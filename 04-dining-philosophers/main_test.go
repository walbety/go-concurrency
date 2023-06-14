package main

import (
	"testing"
	"time"
)

func Test_dineWithVaryingDelays(t *testing.T) {

	var testCases = []struct {
		name  string
		delay time.Duration
	}{
		{
			name:  "zero delay",
			delay: time.Second * 0,
		},
		{
			name:  "quarter second delay",
			delay: time.Millisecond * 250,
		},
		{
			name:  "half second delay",
			delay: time.Millisecond * 500,
		},
	}

	for _, e := range testCases {
		orderFinished = []string{}

		eatTime, thinkTime, sleepTime = e.delay, e.delay, e.delay

		dine()

		if len(orderFinished) != 5 {
			t.Errorf("%s case: incorrect length of slice; expected 5 but got %d", e.name, len(orderFinished))
		}
	}
}
