package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Account struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		Firstname: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(10000000000)),
		CreatedAt: time.Now().UTC(),
	}
}
