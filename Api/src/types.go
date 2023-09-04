package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewAccount(firstName, email, passowrd string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		Email:     email,
		Password:  passowrd,
	}
}
