package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiHandler) handleGetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	var account *Account
	s.db.First(&account, "id = ?", id)
	return s.WriteJson(w, http.StatusOK, account)
}

func (s *ApiHandler) handleGetAllAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
	account := s.NewAccount(createAccount.FirstName, createAccount.Email, createAccount.Password)
	s.db.Create(account)
	return s.WriteJson(w, http.StatusOK, createAccount)
}

func (s *ApiHandler) handleDeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	s.db.Delete(&Account{}, "id = ?", id)
	return s.WriteJson(w, http.StatusOK, "successfully deleted")
}
