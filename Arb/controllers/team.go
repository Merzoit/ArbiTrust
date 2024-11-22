package controllers

import (
	"arb/constants"
	"arb/repositories"
	"arb/structures"
	"encoding/json"
	"log"
	"net/http"
)

type TeamController struct {
	teamRepo repositories.TeamRepository
}

func NewTeamController(repo repositories.TeamRepository) *TeamController {
	return &TeamController{teamRepo: repo}
}

func (tc *TeamController) CreateTeam(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallCreateTeam)
	var team structures.Team
	err := json.NewDecoder(r.Body).Decode(&team)

	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDecodingRequestBody, err)
		http.Error(w, constants.ErrInvalidInput, http.StatusBadRequest)
		return
	}

	if err := tc.teamRepo.AddTeam(&team); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorAddingTeam, err)
		http.Error(w, constants.ErrAddingTeam, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(team)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
	}

	log.Printf("CONTROLLER: "+constants.LogTeamCreateSuccessfully, team.Name)
}

func (tc *TeamController) GetTeamById(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallGetTeamByID)

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	team, err := tc.teamRepo.GetTeamById(id)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorFetchingTeam, err)
		http.Error(w, constants.ErrFetchingTeams, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(team)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogTeamFetchSuccessfully, team.Name)
}

func (tc *TeamController) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallUpdateTeam)

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf(constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	var team structures.Team
	err = json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDecodingRequestBody, err)
		http.Error(w, constants.ErrInvalidInput, http.StatusBadRequest)
		return
	}
	team.ID = id

	if err := tc.teamRepo.UpdateTeam(&team); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorUpdatingTeam, err)
		http.Error(w, constants.ErrUpdatingTeam, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(team)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogTeamUpdateSuccessfully, id)
}

func (tc *TeamController) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallDeleteTeam)

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	if err := tc.teamRepo.DeleteTeam(id); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDeletingTeam, err)
		http.Error(w, constants.ErrDeleteTeam, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("CONTROLLER: "+constants.LogTeamDeleteSuccessfully, id)
}

func (tc *TeamController) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallGetAllTeams)

	teams, err := tc.teamRepo.GetAllTeams()
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorFetchingTeam, err)
		http.Error(w, constants.ErrFetchingTeams, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(teams)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogFetchTeamCount, len(teams))
}
