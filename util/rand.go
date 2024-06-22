package util

import "math/rand"

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
	return rand.Intn(rr.Max-rr.Min+1) + rr.Min
}
