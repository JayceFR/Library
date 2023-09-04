package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiHandler struct {
	db gorm.DB
}

type Account struct {
	ID        int    `gorm:"primarykey" json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func New() *ApiHandler {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Account{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return &ApiHandler{
		db: *db,
	}
}

func (s *ApiHandler) NewAccount(id int, firstName, email, passowrd string) *Account {
	return &Account{
		ID:        id,
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
	ctx := context.Background()
	if r.Method == "GET" {
		return s.handleGetAccount(ctx, w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(ctx, w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(ctx, w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}
