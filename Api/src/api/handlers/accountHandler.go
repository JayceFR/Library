package api

import (
	"context"
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
	fmt.Println(len(response))

	return s.WriteJson(w, http.StatusOK, response)
}

func (s *ApiHandler) handleCreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiHandler) handleDeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
