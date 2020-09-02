package cloudbucket

import (
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/dgryski/trifles/uuid"
	"github.com/factly/bindu-server/config"
	"github.com/factly/x/renderx"
	"google.golang.org/api/option"
)

var (
	storageClient *storage.Client
)

//this function returns the filename(to save in database) of the saved file or an error if it occurs
func fileUpload(r *http.Request) (string, error) {

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

	file, _, err := r.FormFile("file")

	if err != nil {
		return "", err
	}

	defer file.Close() //close the file when we finish

	fileName := uuid.UUIDv4()

	sw := storageClient.Bucket(bucket).Object(data.Details.Path + fileName + ".jpg").NewWriter(ctx)

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

	return u.EscapedPath(), nil

}

// UploadToGCS - file upload to GCS
func UploadToGCS(w http.ResponseWriter, r *http.Request) {

	file, err := fileUpload(r)

	if err != nil {
		renderx.JSON(w, http.StatusOK, err)
		return
	}

	renderx.JSON(w, http.StatusOK, file)
}
