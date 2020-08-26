package chart

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestChartDelete(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(Routes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid chart id", func(t *testing.T) {

		e.DELETE(path).
			WithPath("chart_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)

	})

	t.Run("chart record not found", func(t *testing.T) {

		mock.ExpectQuery(selectQuery).
			WithArgs(100, 1).
			WillReturnRows(sqlmock.NewRows(chartColumns))

		e.DELETE(path).
			WithPath("chart_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("chart record deleted", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(chartColumns).
				AddRow(1, time.Now(), time.Now(), nil, data["title"], data["slug"], byteDescriptionData,
					data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))

		mock.ExpectBegin()
		mock.ExpectExec(deleteQuery).
			WithArgs(test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.DELETE(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
	})

}
