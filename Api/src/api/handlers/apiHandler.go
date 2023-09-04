package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiHandler struct {
	db gorm.DB
}

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func New() *ApiHandler {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return &ApiHandler{
		db: *db,
	}
}

func (s *ApiHandler) NewAccount(firstName, email, passowrd string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		Email:     email,
		Password:  passowrd,
	}
}

func (s *ApiHandler) WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiHandler) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *ApiHandler) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := s.NewAccount("Jayce", "jeffyjany31@gmail.com", "jefjan")
	return s.WriteJson(w, http.StatusOK, account)
}

func (s *ApiHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiHandler) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
