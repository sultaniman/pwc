package util

import (
	"math/rand"
	"time"
)

type RandRange struct {
	Min int
	Max int
}

func NewRandRange(min int, max int) *RandRange {
	return &RandRange{
		Min: min,
		Max: max,
	}
}

func (rr *RandRange) Next() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(rr.Max-rr.Min+1) + rr.Min
}
