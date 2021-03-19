package chart

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/minio"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

var res = map[string]interface{}{
	"title": "Pie",
	"slug":  "pie",
	"description": `{
		"data": [
			{
			"type": "articles",
			"id": "3",
			"attributes": {
				"title": "JSON:API paints my bikeshed!",
				"body": "The shortest article. Ever.",
				"created": "2015-05-22T14:56:29.000Z",
				"updated": "2015-05-22T14:56:28.000Z"
			}
			}
		]
		}`,
	"data_url": "http://data.com/crime?page[number]=3&page[size]=1",
	"config": `{
		"links": {
			"self": "http://example.com/articles?page[number]=3&page[size]=1",
			"first": "http://example.com/articles?page[number]=1&page[size]=1",
			"prev": "http://example.com/articles?page[number]=2&page[size]=1",
			"next": "http://example.com/articles?page[number]=4&page[size]=1",
			"last": "http://example.com/articles?page[number]=13&page[size]=1"
		  }
	}`,
	"status":         "available",
	"published_date": time.Time{},
}

func TestChartCreate(t *testing.T) {

	mock := test.SetupMockDB()

	test.MockServers()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("cannot decode chart", func(t *testing.T) {
		test.CheckSpace(mock)

		e.POST(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("Unprocessable chart", func(t *testing.T) {

		test.CheckSpace(mock)
		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create chart", func(t *testing.T) {
		test.CheckSpace(mock)
		mock.ExpectBegin()

		mock.ExpectQuery(`INSERT INTO "bi_medium"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		slugCheckMock(mock)

		tagQueryMock(mock)

		categoryQueryMock(mock)

		chartInsertMock(mock)

		mock.ExpectQuery(selectQuery).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow("1", time.Now(), time.Now(), nil, 1, 1, data["title"], data["slug"], byteDescriptionData,
					data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))

		chartPreloadMock(mock)

		mock.ExpectCommit()

		result := e.POST(basePath).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(res)

		validateAssociations(result)
		test.ExpectationsMet(t, mock)

	})

	t.Run("create chart with slug is empty", func(t *testing.T) {
		test.CheckSpace(mock)
		mock.ExpectBegin()

		mock.ExpectQuery(`INSERT INTO "bi_medium"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		slugCheckMock(mock)

		tagQueryMock(mock)

		categoryQueryMock(mock)

		chartInsertMock(mock)

		mock.ExpectQuery(selectQuery).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, data["title"], data["slug"], byteDescriptionData,
					data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))

		chartPreloadMock(mock)
		mock.ExpectCommit()

		result := e.POST(basePath).
			WithHeaders(headers).
			WithJSON(dataWithoutSlug).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(res)

		validateAssociations(result)
		test.ExpectationsMet(t, mock)
	})

	t.Run("when uploading returns error", func(t *testing.T) {
		test.CheckSpace(mock)
		minio.Upload = func(r *http.Request, image string) (string, error) {
			return "", errors.New("some error")
		}

		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusInternalServerError)

		test.ExpectationsMet(t, mock)
	})

}
