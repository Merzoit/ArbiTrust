package controllers

import (
	"arb/repositories"
	"arb/structures"
	"encoding/json"
	"net/http"
)

type TeamController struct {
	teamRepo repositories.TeamRepository
}

func NewTeamController(repo repositories.TeamRepository) *TeamController {
	return &TeamController{teamRepo: repo}
}

func (tc *TeamController) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team structures.Team
	err := json.NewDecoder(r.Body).Decode(&team)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := tc.teamRepo.AddTeam(&team); err != nil {
		http.Error(w, "Failed to add team", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

func (tc *TeamController) GetTeamById(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	team, err := tc.teamRepo.GetTeamById(id)
	if err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(team)
}

func (tc *TeamController) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var team structures.Team
	err = json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	team.ID = id

	if err := tc.teamRepo.UpdateTeam(&team); err != nil {
		http.Error(w, "Failed to update team", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(team)
}

func (tc *TeamController) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := tc.teamRepo.DeleteTeam(id); err != nil {
		http.Error(w, "Failed to delete team", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (tc *TeamController) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := tc.teamRepo.GetAllTeams()
	if err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
