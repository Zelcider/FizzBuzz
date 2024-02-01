package command

import (
	"fizzbuzz/api/internal/domain/entities"
	"fizzbuzz/api/internal/domain/repositories"
)

type CreateFizzBuzzCommand struct {
	Limit int64
	Int1  int64
	Str1  string
	Int2  int64
	Str2  string
	Repo  repositories.IStat
}

func (command *CreateFizzBuzzCommand) Execute() (string, error) {
	// Add Stats
	command.Repo.AddRequest(command.ToStatRequest())

	// Create FizzBuzz
	fizzbuzz, err := entities.CreateFizzBuzz(command.Limit, command.Int1, command.Str1, command.Int2, command.Str2)
	if err != nil {
		return "", err
	}

	// Compute FizzBuzz
	return fizzbuzz.Compute(), nil
}
