package entities

import (
	"errors"
	"strconv"
	"strings"
)

type FizzBuzz struct {
	Limit int64
	Int1  int64
	Str1  string
	Int2  int64
	Str2  string
}

func CreateFizzBuzz(
	Limit int64,
	Int1 int64,
	Str1 string,
	Int2 int64,
	Str2 string,
) (FizzBuzz, error) {
	if Int1 < 1 {
		return FizzBuzz{}, errors.New("Invalid value  (" + strconv.FormatInt(Int1, 10) + ") for Int1, non nullable positive integer expected")
	}
	if Int2 < 1 {
		return FizzBuzz{}, errors.New("Invalid value (" + strconv.FormatInt(Int2, 10) + ") for Int2, non nullable positive integer expected")
	}
	if Str1 == "" {
		return FizzBuzz{}, errors.New("Invalid value (" + Str1 + ") for Str1, non-empty string expected")
	}
	if Str2 == "" {
		return FizzBuzz{}, errors.New("Invalid value (" + Str2 + ")for Str2, non-empty string expected")
	}
	if Limit < 1 {
		return FizzBuzz{}, errors.New("Invalid value (" + strconv.FormatInt(Limit, 10) + ") for Limit, non nullable positive integer expected")
	}

	return FizzBuzz{
		Limit: Limit,
		Int1:  Int1,
		Str1:  Str1,
		Int2:  Int2,
		Str2:  Str2,
	}, nil
}

const start int64 = 1

func (fizzBuzz *FizzBuzz) Compute() string {
	var output strings.Builder
	separator := ""
	for i := start; i <= fizzBuzz.Limit; i++ {

		if i > start {
			separator = ","
		}

		if i%fizzBuzz.Int1 == 0 && i%fizzBuzz.Int2 == 0 {
			output.WriteString(separator + fizzBuzz.Str1 + fizzBuzz.Str2)
			continue
		}

		if i%fizzBuzz.Int1 == 0 {
			output.WriteString(separator + fizzBuzz.Str1)
			continue
		}

		if i%fizzBuzz.Int2 == 0 {
			output.WriteString(separator + fizzBuzz.Str2)
			continue
		}

		output.WriteString(separator + strconv.FormatInt(i, 10))

	}
	return output.String()
}
