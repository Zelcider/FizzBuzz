package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_FizzBuzz_Api_Should_Return_400(t *testing.T) {
	//Arrange
	serverApi := Setup()
	request, _ := http.NewRequest("POST", "/api/v1/fizzbuzz", bytes.NewBuffer(nil))

	//Act
	response := serverApi.Call(request)

	//Assert
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func Test_FizzBuzz_Api_Should_Return_200(t *testing.T) {
	//Arrange
	serverApi := Setup()
	createFizzBuzz, _ := json.Marshal(map[string]interface{}{"limit": 15, "int1": 3, "str1": "fizz", "int2": 5, "str2": "buzz"})

	request, _ := http.NewRequest("POST", "/api/v1/fizzbuzz", bytes.NewBuffer(createFizzBuzz))

	//Act
	response := serverApi.Call(request)

	//Assert
	assert.Equal(t, http.StatusCreated, response.Code)
}
