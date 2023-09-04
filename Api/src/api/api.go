package api

import (
	"encoding/json"
	"log"
	"net/http"

	handlers "github.com/JayceFR/library/src/api/handlers"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			//handle the error
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	ApiHandler := handlers.New()
	log.Println("Api running on port :", s.listenAddr)
	router.HandleFunc("/account", makeHttpHandleFunc(ApiHandler.HandleAccount))
	http.ListenAndServe(s.listenAddr, router)
}
