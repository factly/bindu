package theme

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestThemeUpdate(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	var updatedTheme = map[string]interface{}{
		"name": "Dark theme",
		"config": `{"image": { 
			"src": "Images/Sun.png",
			"name": "sun2",
			"hOffset": 25,
			"vOffset": 20,
			"alignment": "left"
		}}`,
	}

	updatedByteData, _ := json.Marshal(updatedTheme["config"])

	t.Run("invalid theme id", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("theme_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
	})

	t.Run("cannot decode theme", func(t *testing.T) {

		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("theme_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("Unprocessable theme", func(t *testing.T) {

		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("theme_id", 1).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("theme record not found", func(t *testing.T) {
		test.CheckSpace(mock)
		recordNotFoundMock(mock)

		e.PUT(path).
			WithPath("theme_id", "100").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update theme", func(t *testing.T) {

		test.CheckSpace(mock)
		SelectMock(mock)

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_theme\"`).
			WithArgs(test.AnyTime{}, 1, updatedTheme["name"], updatedByteData, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, updatedTheme["name"], updatedByteData, 1))

		e.PUT(path).
			WithPath("theme_id", 1).
			WithHeaders(headers).
			WithJSON(updatedTheme).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedTheme)

	})

}
