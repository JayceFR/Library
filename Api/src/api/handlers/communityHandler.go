package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateCommunity struct {
	CommunityName string `json:"community_name"`
}

func (s *ApiHandler) handleCreateCommunity(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	createCommunity := CreateCommunity{}
	err := json.NewDecoder(r.Body).Decode(&createCommunity)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(createCommunity.CommunityName)
	checkCommunity := Community{}
	s.db.First(&checkCommunity, "community_name = ?", createCommunity.CommunityName)
	if checkCommunity.ID.String() == null_uuid {
		fmt.Println("Community name is unique ")
		community := s.NewCommunity(createCommunity.CommunityName)
		s.db.Create(community)
		return s.WriteJson(w, http.StatusOK, community)
	} else {
		return s.WriteJson(w, http.StatusBadRequest, "Found another community with the same name")
	}
}
