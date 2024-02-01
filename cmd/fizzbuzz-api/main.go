package main

import (
	"fizzbuzz/api/internal/infrastructure/api"
	"fizzbuzz/api/internal/infrastructure/db"
)

func main() {
	var repositories = api.Repositories{Stat: db.NewStatInMemory()}
	router := api.Setup(repositories)
	router.Run()
}
