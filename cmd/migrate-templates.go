package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	minioutil "github.com/factly/bindu-server/util/minio"
	"github.com/factly/x/requestx"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	TemplatesPath string            = "./templates"
	headers       map[string]string = map[string]string{
		"X-Space": viper.GetString("migration_space"),
		"X-User":  viper.GetString("migration_user"),
	}
)

func init() {
	rootCmd.AddCommand(migrateTemplatesCmd)

	config.SetupVars()

	minioutil.SetupClient()

	TemplatesPath = "./templates"
	headers = map[string]string{
		"X-Space": fmt.Sprint(viper.GetInt("migration_space")),
		"X-User":  fmt.Sprint(viper.GetInt("migration_user")),
	}
}

var migrateTemplatesCmd = &cobra.Command{
	Use:   "migrate-templates",
	Short: "Apply migrations for templates data for bindu-server.",
	Run: func(cmd *cobra.Command, args []string) {
		err := MigrateTemplate()
		if err != nil {
			log.Fatal(err)
		}
	},
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

	_, migrated := CategoriesMigrated(categories)

	if !migrated {
		category_map := make(map[string]uint)
		for _, category_name := range categories {
			category := model.Category{
				Name: category_name,
			}

			resp, err := requestx.Request("POST", "http://localhost:8000/categories", category, headers)
			if err != nil {
				return err
			}

			if resp.StatusCode != http.StatusCreated {
				return errors.New("could not create category " + category_name)
			}
			defer resp.Body.Close()

			respCategory := model.Category{}
			if err := json.NewDecoder(resp.Body).Decode(&respCategory); err != nil {
				return err
			}

			category_map[respCategory.Name] = respCategory.ID
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
				var properties []map[string]interface{}
				propertiesFile, err := os.Open(fmt.Sprint(filepath, "/properties.json"))
				if err != nil {
					return err
				}
				defer propertiesFile.Close()

				bytes, _ := ioutil.ReadAll(propertiesFile)
				err = json.Unmarshal(bytes, &properties)
				if err != nil {
					return err
				}

				// fetching spec
				var spec map[string]interface{}
				specFile, err := os.Open(fmt.Sprint(filepath, "/spec.json"))
				if err != nil {
					return err
				}
				defer specFile.Close()

				bytes, _ = ioutil.ReadAll(specFile)
				err = json.Unmarshal(bytes, &spec)
				if err != nil {
					return err
				}

				mediumID, err := CreateMedium(filepath, fmt.Sprint(chart_name, ".png"), "thumbnail.png")
				if err != nil {
					return err
				}
				fmt.Println(`created medium`, chart_name)

				templateBody := map[string]interface{}{
					"category_id": category_map[category_name],
					"medium_id":   mediumID,
					"properties":  properties,
					"spec":        spec,
					"title":       chart_name,
					"slug":        strings.ToLower(chart_name),
				}

				resp, err := requestx.Request("POST", "http://localhost:8000/templates", templateBody, headers)
				if err != nil {
					return err
				}

				if resp.StatusCode != http.StatusCreated {
					return errors.New(`cannot create template ` + chart_name)
				} else {
					fmt.Println("template " + chart_name + " created")
				}
			}
		}
	} else {
		fmt.Println("migrations done...")
	}

	return nil
}

func CreateMedium(path, chartName, filename string) (uint, error) {
	info, err := minioutil.Client.FPutObject(context.Background(), "dega", fmt.Sprint("bindu/", chartName), fmt.Sprint(path, "/", filename), minio.PutObjectOptions{})
	if err != nil {
		return 0, err
	}

	mediumBody := map[string]interface{}{
		"name": chartName,
		"url": map[string]interface{}{
			"raw": fmt.Sprint("http://", viper.GetString("minio_url"), "/dega/bindu/", chartName),
		},
		"file_size": info.Size,
	}

	resp, err := requestx.Request("POST", "http://localhost:8000/media", mediumBody, headers)
	if err != nil {
		return 0, err
	}

	var respMedium model.Medium
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respMedium)
	if err != nil {
		return 0, err
	}

	return respMedium.ID, nil
}

func CategoriesMigrated(categoryNames []string) (map[string]uint, bool) {

	resp, err := requestx.Request("GET", "http://localhost:8000/categories", nil, headers)
	if err != nil {
		return nil, false
	}

	type catPage struct {
		Nodes []model.Category `json:"nodes,omitempty"`
		Total int              `json:"total,omitempty"`
	}

	categoriesPaiganation := catPage{}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&categoriesPaiganation); err != nil {
		return nil, false
	}

	categories := categoriesPaiganation.Nodes
	categoryMap := make(map[string]uint)
	for _, cat := range categories {
		categoryMap[cat.Name] = cat.ID
	}

	for _, category := range categoryNames {
		if _, found := categoryMap[category]; !found {
			return nil, false
		}
	}
	return categoryMap, true
}
