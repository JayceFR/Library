package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *ApiHandler) HandleLoginAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	loginAccount := LoginAccount{}
	err := json.NewDecoder(r.Body).Decode(&loginAccount)
	if err != nil {
		fmt.Println(err.Error())
	}
	var check_account Account
	s.db.First(&check_account, "email = ? AND password = ?", loginAccount.Email, loginAccount.Password)
	return s.WriteJson(w, http.StatusOK, check_account)
}
