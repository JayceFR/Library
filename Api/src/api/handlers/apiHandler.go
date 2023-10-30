package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiHandler struct {
	db gorm.DB
}

type Account struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
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

func (s *ApiHandler) NewAccount(firstName string, email string, passowrd []byte) *Account {
	id := uuid.New()
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
		return s.handleGetAllAccount(ctx, w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(ctx, w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiHandler) HandleSpecificAccount(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	if r.Method == "GET" {
		return s.handleGetAccount(ctx, w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(ctx, w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiHandler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	if r.Method == "POST" {
		return s.HandleLoginAccount(ctx, w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
