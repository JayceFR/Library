package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *ApiHandler) handleGetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// account := s.NewAccount("Jayce", "jeffyjany31@gmail.com", "jefjan")
	// s.db.Create(account)
	// var get_account Account
	// s.db.First(&get_account, "first_name = ?", "Jayce")
	response, err := s.GetAllAcounts(ctx, s.db)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s.WriteJson(w, http.StatusOK, response)
}

type CreateAccount struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (s *ApiHandler) handleCreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	createAccount := CreateAccount{}
	err := json.NewDecoder(r.Body).Decode(&createAccount)
	if err != nil {
		fmt.Println(err.Error())
	}
	response, err := s.GetAllAcounts(ctx, s.db)
	if err != nil {
		fmt.Println(err.Error())
	}
	account := s.NewAccount(len(response), createAccount.FirstName, createAccount.Email, createAccount.Password)
	s.db.Create(account)
	return s.WriteJson(w, http.StatusOK, createAccount)
}

func (s *ApiHandler) handleDeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
