package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		Firstname: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(10000000000)),
	}
}
