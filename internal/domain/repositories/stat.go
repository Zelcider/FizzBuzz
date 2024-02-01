package repositories

import (
	"fizzbuzz/api/internal/domain/entities"
)

type IStat interface {
	GetMostFrequentRequest() entities.Stat
	AddRequest(request entities.StatRequest)
}
