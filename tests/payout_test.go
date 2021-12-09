package tests

/*
These are integration tests.
*/

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"takeHomeTest/controllers"
	"takeHomeTest/errorHandlers"
	"takeHomeTest/repository"
	"takeHomeTest/service"
	"testing"
)

var (
	repo         repository.Repository
	serv         service.Service
	controller   controllers.Controller
	errorHandler errorHandlers.ErrorHandlers
	GoodJson     = `[
    {
        "name": "Top",
        "price_amount":10000,
        "price_currency":"USD",
        "seller_reference": 1
    }]`

	BadJson = `[
    {
        "price_amount":10000,
        "price_currency":"USD",
        "seller_reference": 1
    }]`
)

func initialize(db *sql.DB) {
	errorHandler = errorHandlers.NewErrorHandler()
	repo = repository.NewRepository(db, errorHandler)
	serv = service.NewService(repo, errorHandler)
	controller = controllers.NewController(serv, errorHandler)
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.POST("/createPayout", controller.CreatePayout)
	return r
}

func TestCreatePayout_whenRequestBodyGiven_ThenOkStatus(t *testing.T) {

	//given
	db, mock := NewMock()
	initialize(db)
	defer db.Close()
	router := SetupRouter()
	w := httptest.NewRecorder()

	//when
	req, err := http.NewRequest("POST", "/createPayout", strings.NewReader(GoodJson))

	//then
	assert.NoError(t, err)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Payout").WithArgs(1, 10000, "USD").WillReturnResult(sqlmock.NewResult(2, 2))
	mock.ExpectCommit()
	router.ServeHTTP(w, req)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There are unfulfilled expections: %s", err)
	}
	response := w.Body.String()
	assert.NotNil(t, response)
	assert.Contains(t, response, "10000")
	assert.Contains(t, response, "\"no_of_transactions\":1")
	assert.Contains(t, response, "USD")
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestCreatePayout_whenBadRequestBodyGiven_ThenBadRequest(t *testing.T) {

	//given
	db, _ := NewMock()
	initialize(db)
	defer db.Close()
	router := SetupRouter()
	w := httptest.NewRecorder()

	//when
	req, err := http.NewRequest("POST", "/createPayout", strings.NewReader(BadJson))

	//then
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
	assert.Contains(t, w.Body.String(), "less than min")
}
