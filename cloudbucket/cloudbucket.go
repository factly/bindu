package cloudbucket

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/factly/bindu-server/config"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

var (
	storageClient *storage.Client
)

//FileUpload function returns the filename(to save in database) of the saved file or an error if it occurs
func FileUpload(r *http.Request, file io.Reader) (string, error) {

	data, jsonByte, err := config.GetBucketConfig()

	if err != nil {
		return "", err
	}

	bucket := data.Details.BucketName

	ctx := r.Context()

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsJSON(jsonByte))
	if err != nil {
		return "", err
	}

	r.ParseMultipartForm(int64(data.Details.Limit))

	fileName := uuid.New()

	sw := storageClient.Bucket(bucket).Object(fmt.Sprint(data.Details.Path, fileName, ".jpg")).NewWriter(ctx)

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
