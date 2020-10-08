package cloudbucket

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

var (
	storageClient *storage.Client
)

//FileUpload function returns the filename(to save in database) of the saved file or an error if it occurs
func FileUpload(r *http.Request, file io.Reader) (string, error) {

	jsonByte, err := json.Marshal(viper.GetStringMapString("bucket.keys"))

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	bucket := viper.GetString("bucket.details.bucket_name")

	ctx := r.Context()

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsJSON(jsonByte))
	if err != nil {
		return "", err
	}

	r.ParseMultipartForm(int64(viper.GetInt("bucket.details.limit")))

	fileName := uuid.New()

	sw := storageClient.Bucket(bucket).Object(fmt.Sprint(viper.GetString("bucket.details.path"), fileName, ".jpg")).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return "", err
	}

	path := "https://storage.googleapis.com" + u.EscapedPath()

	return path, nil

}
