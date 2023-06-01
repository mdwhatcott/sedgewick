package ch3p4

import (
	"testing"
	"time"
)

var m = 367
var R = 10

func DateHash(t time.Time) int {
	day := t.Day()
	month := int(t.Month())
	year := t.Year()
	return (((day*R+month)%m)*R + year) % m
}

func TestDateHash(t *testing.T) {
	frequencies := make(map[int]int) // map[hash]hits
	now := time.Now()
	date := time.Unix(0, 0)
	for date.Before(now) {
		date = date.AddDate(0, 0, 1)
		frequencies[DateHash(date)]++
	}
	var max, min = 0, 1000
	for key, val := range frequencies {
		t.Log(key, val)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	t.Log("MAX:", max)
	t.Log("MIN:", min)
}
