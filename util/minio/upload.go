package minio

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/minio/minio-go/v7"

	"github.com/spf13/viper"
)

// Upload uploads the image to minio
var Upload = func(r *http.Request, image string) (string, error) {

	idx := strings.Index(image, ";base64,")
	if idx < 0 {
		return "", errors.New("invalid image")
	}
	imageType := image[11:idx]

	unbased, err := base64.StdEncoding.DecodeString(image[idx+8:])
	if err != nil {
		return "", err
	}
	file := bytes.NewReader(unbased)

	fileName := fmt.Sprint(uuid.New(), ".", imageType)

	info, err := Client.PutObject(r.Context(), viper.GetString("minio.bucket"), fileName, file, -1, minio.PutObjectOptions{
		ContentType: "image/" + imageType,
	})
	return info.Location, err
}
