package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	minioutil "github.com/factly/bindu-server/util/minio"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/slugx"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	TemplatesPath string = "./templates"
	SpaceID       uint   = 0
)

func init() {
	rootCmd.AddCommand(migrateTemplatesCmd)

	config.SetupVars()
	config.SetupDB()

	minioutil.SetupClient()

	TemplatesPath = "./templates"

}

var migrateTemplatesCmd = &cobra.Command{
	Use:   "migrate-templates",
	Short: "Apply migrations for templates data for bindu-server.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		SpaceID, err = GetSuperOrgSpace()
		if err != nil {
			log.Fatal(err.Error())
		}

		err = MigrateTemplate()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func GetSuperOrgSpace() (uint, error) {
	soID, err := middlewarex.GetSuperOrganisationID("bindu")
	if err != nil {
		return 0, err
	}

	spaces := make([]model.Space, 0)
	err = config.DB.Model(&model.Space{}).Where(&model.Space{
		OrganisationID: soID,
	}).Find(&spaces).Error
	if err != nil {
		return 0, err
	}

	for _, space := range spaces {
		if space.Name == viper.GetString("super_space_name") {
			return space.ID, nil
		}
	}
	return 0, errors.New("cannot get super space")
}

func MigrateTemplate() error {
	categories_paths := make([]string, 0)
	categories := make([]string, 0)

	files, err := ioutil.ReadDir(TemplatesPath)
	if err != nil {
		return err
	}

	for _, each := range files {
		categories = append(categories, each.Name())
		categories_paths = append(categories_paths, fmt.Sprint(TemplatesPath, "/", each.Name()))
	}

	category_map := make(map[string]uint)
	presentCategories := make([]model.Category, 0)
	// fetch all categories
	config.DB.Model(&model.Category{}).Find(&presentCategories)

	for _, cat := range presentCategories {
		category_map[cat.Name] = cat.ID
	}

	// Adds categories if any new found
	for _, category_name := range categories {
		if _, found := category_map[category_name]; !found {
			category := model.Category{
				Name:          category_name,
				IsForTemplate: true,
				SpaceID:       SpaceID,
			}

			if err = config.DB.Create(&category).Error; err != nil {
				return err
			}
			category_map[category.Name] = category.ID
		}
	}

	for _, cat_path := range categories_paths {
		files, err := ioutil.ReadDir(cat_path)
		if err != nil {
			return err
		}

		fmt.Println("Processing files in " + cat_path)

		for _, file := range files {
			filepath := fmt.Sprint(cat_path, "/", file.Name())
			category_name := strings.Split(cat_path, "/")[2]
			chart_name := file.Name()
			fmt.Println("Processing ", filepath)

			// fetching properties
			propertiesFile, err := os.Open(fmt.Sprint(filepath, "/properties.json"))
			if err != nil {
				return err
			}
			defer propertiesFile.Close()

			propertiesBytes, err := ioutil.ReadAll(propertiesFile)
			if err != nil {
				return err
			}

			// fetching spec
			specFile, err := os.Open(fmt.Sprint(filepath, "/spec.json"))
			if err != nil {
				return err
			}
			defer specFile.Close()

			specBytes, err := ioutil.ReadAll(specFile)
			if err != nil {
				return err
			}

			var spec map[string]interface{}
			err = json.Unmarshal(specBytes, &spec)
			if err != nil {
				return err
			}

			mode := spec["mode"].(string)

			mediumID, err := CreateMedium(filepath, fmt.Sprint(chart_name, ".png"), "thumbnail.png")
			if err != nil {
				return err
			}

			template := model.Template{
				CategoryID: category_map[category_name],
				MediumID:   &mediumID,
				Properties: postgres.Jsonb{
					RawMessage: propertiesBytes,
				},
				Spec: postgres.Jsonb{
					RawMessage: specBytes,
				},
				Title:     strcase.ToCamel(chart_name),
				IsDefault: true,
				Slug:      slugx.Make(fmt.Sprint("bindu ", chart_name)),
				Mode:      mode,
				SpaceID:   SpaceID,
			}

			migratedID, err := ioutil.ReadFile(fmt.Sprint(filepath, "/migrate.out"))
			if err != nil || migratedID == nil {
				// file not found
				template.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
				if err = config.DB.Create(&template).Error; err != nil {
					return err
				} else {
					fmt.Println("template " + chart_name + " created")
				}
				err = ioutil.WriteFile(fmt.Sprint(filepath, "/migrate.out"), []byte(template.ID), 0777)
				if err != nil {
					return err
				}
			} else {
				// file found
				presentTemplate := model.Template{}
				presentTemplate.ID = string(migratedID)

				if err = config.DB.Model(&presentTemplate).Where(&presentTemplate).First(&model.Template{}).Error; err == nil {
					// template found
					if err = config.DB.Model(&presentTemplate).Updates(template).Error; err != nil {
						return err
					} else {
						fmt.Println("template " + chart_name + " updated")
					}
				} else { // template not found
					template.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
					if err = config.DB.Create(&template).Error; err != nil {
						return err
					} else {
						fmt.Println("template " + chart_name + " created")
					}
					err = ioutil.WriteFile(fmt.Sprint(filepath, "/migrate.out"), []byte(template.ID), 0777)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func CreateMedium(path, chartName, filename string) (uint, error) {
	info, err := minioutil.Client.FPutObject(context.Background(), viper.GetString("minio_bucket"), fmt.Sprint("bindu/", chartName), fmt.Sprint(path, "/", filename), minio.PutObjectOptions{})
	if err != nil {
		return 0, err
	}

	urlBytes, _ := json.Marshal(map[string]interface{}{
		"raw": fmt.Sprint("http://", viper.GetString("minio_public_url"), "/", viper.GetString("minio_bucket"), "/bindu/", chartName),
	})

	medium := model.Medium{
		Name: chartName,
		URL: postgres.Jsonb{
			RawMessage: urlBytes,
		},
		FileSize: info.Size,
		SpaceID:  SpaceID,
	}

	if err = config.DB.Model(&model.Medium{}).Where(&medium).First(&medium).Error; err == nil {
		return medium.ID, nil
	} else {
		// create medium
		if err = config.DB.Create(&medium).Error; err != nil {
			return medium.ID, err
		}
		fmt.Println("created medium ", chartName)
	}

	return medium.ID, nil
}
