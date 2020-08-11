package theme

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
)

func TestThemeCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Post("/themes", create)

	var jsonStr = []byte(`
	{
		"name": "Dark theme"
	}`)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Unprocessable theme", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "POST", "/themes", nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("create theme", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/themes", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["name"] != "Dark theme" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Dark theme")
		}

	})

}
