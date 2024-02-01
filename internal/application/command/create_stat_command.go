package command

import (
	"fizzbuzz/api/internal/domain/entities"
	"fizzbuzz/api/internal/domain/repositories"
)

type CreateStatCommand struct {
	Repo    repositories.IStat
	Request CreateFizzBuzzCommand
}

func (command *CreateFizzBuzzCommand) ToStatRequest() entities.StatRequest {
	return entities.StatRequest{
		Limit: command.Limit,
		Int1:  command.Int1,
		Str1:  command.Str1,
		Int2:  command.Int2,
		Str2:  command.Str2,
	}
}

func (command *CreateStatCommand) Execute() error {
	command.Repo.AddRequest(command.Request.ToStatRequest())
	return nil
}
