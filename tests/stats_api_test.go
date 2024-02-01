package tests

import (
	"bytes"
	"fizzbuzz/api/internal/domain/entities"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Stats_Api_Should_Return_Empty_Stats(t *testing.T) {
	//Arrange
	serverApi := Setup()
	request, _ := http.NewRequest("GET", "/api/v1/stats", bytes.NewBuffer(nil))

	//Act
	response := serverApi.Call(request)
	var stats entities.Stat
	json.Unmarshal(response.Body.Bytes(), stats)

	//Assert
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, entities.Stat{}, stats)
}

func Test_Stats_Api_Should_Return_Stats(t *testing.T) {
	//Arrange
	serverApi := Setup()
	CallFizzBuzz(serverApi)
	CallFizzBuzz(serverApi)
	CallFizzBuzz(serverApi)
	request, _ := http.NewRequest("GET", "/api/v1/stats", bytes.NewBuffer(nil))

	//Act
	response := serverApi.Call(request)
	var stats entities.Stat
	err := json.Unmarshal(response.Body.Bytes(), &stats)

	//Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, entities.Stat{
		Request: entities.StatRequest{
			Limit: 15,
			Int1:  3,
			Str1:  "fizz",
			Int2:  5,
			Str2:  "buzz",
		},
		Hits: 3,
	}, stats)
}
