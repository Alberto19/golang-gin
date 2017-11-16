package main

import (
	"time"
)

var customer = map[string]Customer{
	"1": Customer{
		id:        1,
		firstName: "junior",
		lastName:  "farias",
		birthDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
		isDeleted: false,
	},
}

type Customer struct {
	id        uint
	firstName string
	lastName  string
	birthDate time.Time
	isDeleted bool
}
