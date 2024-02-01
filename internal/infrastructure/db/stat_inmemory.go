package db

import (
	"fizzbuzz/api/internal/domain/entities"
	"reflect"
)

type StatInMemory struct {
	store []entities.Stat
}

func NewStatInMemory() *StatInMemory {
	return &StatInMemory{store: make([]entities.Stat, 0)}
}

func (repo *StatInMemory) GetMostFrequentRequest() entities.Stat {
	var maxStat entities.Stat
	for i, e := range repo.store {
		if i == 0 || e.Hits > maxStat.Hits {
			maxStat = e
		}
	}
	return maxStat
}

func (repo *StatInMemory) AddRequest(request entities.StatRequest) {
	found := false
	for i, value := range repo.store {
		if reflect.DeepEqual(value.Request, request) {
			repo.store[i] = entities.Stat{Request: request, Hits: value.Hits + 1}
			found = true
		}
	}

	if found != true {
		repo.store = append(repo.store, entities.Stat{Request: request, Hits: 1})
	}
}
