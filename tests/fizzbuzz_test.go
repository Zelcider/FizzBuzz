package tests

import (
	command "fizzbuzz/api/internal/application/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FizzBuzz_Should_Return_Error_Invalid_Int1(t *testing.T) {
	//Arrange
	createFizzBuzzCommand := command.CreateFizzBuzzCommand{
		Limit: 10,
		Int1:  0,
		Str1:  "test",
		Int2:  2,
		Str2:  "test2",
	}

	//Act
	output, err := createFizzBuzzCommand.Execute()

	//Assert
	assert.EqualError(t, err, "Invalid value for Int1, non nullable positive integer expected")
	assert.Equal(t, output, "")
}

func Test_FizzBuzzShould_Return_Error_Invalid_Int2(t *testing.T) {
	//Arrange
	createFizzBuzzCommand := command.CreateFizzBuzzCommand{
		Limit: 10,
		Int1:  1,
		Str1:  "test",
		Int2:  0,
		Str2:  "test2",
	}

	//Act
	output, err := createFizzBuzzCommand.Execute()

	//Assert
	assert.EqualError(t, err, "Invalid value for Int2, non nullable positive integer expected")
	assert.Equal(t, output, "")
}

func Test_FizzBuzzShould_Return_Error_Invalid_Str1(t *testing.T) {
	//Arrange
	createFizzBuzzCommand := command.CreateFizzBuzzCommand{
		Limit: 10,
		Int1:  1,
		Str1:  "",
		Int2:  1,
		Str2:  "test2",
	}

	//Act
	output, err := createFizzBuzzCommand.Execute()

	//Assert
	assert.EqualError(t, err, "Invalid value for Str1, non-empty string expected")
	assert.Equal(t, output, "")
}

func Test_FizzBuzzShould_Return_Error_Invalid_Str2(t *testing.T) {
	//Arrange
	createFizzBuzzCommand := command.CreateFizzBuzzCommand{
		Limit: 10,
		Int1:  1,
		Str1:  "test",
		Int2:  1,
		Str2:  "",
	}

	//Act
	output, err := createFizzBuzzCommand.Execute()

	//Assert
	assert.EqualError(t, err, "Invalid value for Str2, non-empty string expected")
	assert.Equal(t, output, "")
}

func Test_FizzBuzzShould_Return_Error_Invalid_Limit(t *testing.T) {
	//Arrange
	createFizzBuzzCommand := command.CreateFizzBuzzCommand{
		Limit: 0,
		Int1:  1,
		Str1:  "test",
		Int2:  1,
		Str2:  "test2",
	}

	//Act
	output, err := createFizzBuzzCommand.Execute()

	//Assert
	assert.EqualError(t, err, "Invalid value for Limit, non nullable positive integer expected")
	assert.Equal(t, output, "")
}

func Test_FizzBuzzShould_Return_Output(t *testing.T) {
	//Arrange
	var myTests = []struct {
		command  command.CreateFizzBuzzCommand
		expected string
	}{
		{command: command.CreateFizzBuzzCommand{
			Limit: 10,
			Int1:  2,
			Str1:  "fizz",
			Int2:  3,
			Str2:  "buzz",
		}, expected: "1,fizz,buzz,fizz,5,fizzbuzz,7,fizz,buzz,fizz"},
	}

	for _, tt := range myTests {

		//Act
		output, err := tt.command.Execute()

		//Assert
		assert.Equal(t, nil, err)
		assert.Equal(t, tt.expected, output)
	}
}
