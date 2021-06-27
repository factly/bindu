package organisation

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/permission/organisation"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestOrganisationRequestApprove(t *testing.T) {
	mock := test.SetupMockDB()
	test.SetEnv()

	test.MockServers()
	defer gock.DisableNetworking()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid request id", func(t *testing.T) {
		test.CheckSpace(mock)
		e.POST(approvePath).
			WithPath("request_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
	})

	t.Run("request record not found", func(t *testing.T) {
		test.CheckSpace(mock)
		mock.ExpectQuery(selectQuery).
			WithArgs("pending", 1).
			WillReturnRows(sqlmock.NewRows(Columns))

		e.POST(approvePath).
			WithPath("request_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("organisation permission for the request already exist", func(t *testing.T) {
		test.CheckSpace(mock)
		SelectQuery(mock, "pending", 1)

		organisation.SelectQuery(mock)

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_organisation_permission\"`).
			WithArgs(test.AnyTime{}, 1, Data["organisation_id"], Data["spaces"], 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		organisation.SelectQuery(mock)

		mock.ExpectExec(`UPDATE \"bi_organisation_permission_request\"`).
			WithArgs(test.AnyTime{}, 1, "approved", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.POST(approvePath).
			WithPath("request_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})

	t.Run("approve request and create organisation permission", func(t *testing.T) {
		test.CheckSpace(mock)
		SelectQuery(mock, "pending", 1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_organisation_permission"`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "spaces"}))

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_organisation_permission"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, Data["organisation_id"], Data["spaces"]).
			WillReturnRows(sqlmock.
				NewRows([]string{"id"}).
				AddRow(1))

		mock.ExpectExec(`UPDATE \"bi_organisation_permission_request\"`).
			WithArgs(test.AnyTime{}, 1, "approved", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.POST(approvePath).
			WithPath("request_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})

}
