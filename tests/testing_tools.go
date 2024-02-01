package tests

import (
	"bytes"
	"encoding/json"
	"fizzbuzz/api/internal/infrastructure/api"
	"fizzbuzz/api/internal/infrastructure/db"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

type TestApi interface {
	Setup() TestApiResult
	Run() error
	Call(r http.Request) (http.Response, error)
}

type TestApiResult struct {
	api *api.Result
}

func Setup() *TestApiResult {
	return &TestApiResult{api: api.Setup(api.Repositories{Stat: db.NewStatInMemory()})}
}

func (testApi *TestApiResult) Run() error {
	return testApi.api.Run()
}

func (testApi *TestApiResult) Call(request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	testApi.api.Context.ServeHTTP(response, request)
	return response
}

func CallFizzBuzz(serverApi *TestApiResult) {
	createFizzBuzz, _ := json.Marshal(map[string]interface{}{"limit": 15, "int1": 3, "str1": "fizz", "int2": 5, "str2": "buzz"})
	requestFizzbuzz, _ := http.NewRequest("POST", "/api/v1/fizzbuzz", bytes.NewBuffer(createFizzBuzz))
	serverApi.Call(requestFizzbuzz)
}
