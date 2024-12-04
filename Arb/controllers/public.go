package controllers

import (
	"arb/constants"
	"arb/repositories"
	"arb/structures"
	"encoding/json"
	"log"
	"net/http"
)

type PublicController struct {
	repo repositories.PublicRepository
}

func NewPublicController(repo repositories.PublicRepository) *PublicController {
	return &PublicController{repo: repo}
}

func (pc *PublicController) CreatePublic(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: CreatePublic API called")

	var public structures.Public
	if err := json.NewDecoder(r.Body).Decode(&public); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDecodingRequestBody, err)
		http.Error(w, constants.ErrInvalidInput, http.StatusBadRequest)
		return
	}

	if err := pc.repo.CreatePublic(&public); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorAddingPublic, err)
		http.Error(w, constants.LogErrorAddingPublic, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(public)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogPublicCreateSuccessfully, public.Name)
}

func (pc *PublicController) GetPublicByID(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: GetPublicByID API called")

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	public, err := pc.repo.GetPublicByID(id)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorFetchingPublic, err)
		http.Error(w, constants.ErrFetchingPublic, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(public)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogPublicCreateSuccessfully, public.Name)
}

func (pc *PublicController) GetAllPublics(w http.ResponseWriter, r *http.Request) {
	log.Printf("CONTROLLER: GetAllPublics API called")

	publics, err := pc.repo.GetAllPublics()
	if err != nil {
		log.Printf("CONTROLLER:"+constants.LogErrorFetchingAllPublics, err)
		http.Error(w, constants.ErrFetchingPublic, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(publics)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	log.Println("CONTROLLER: Successfully fetched publics")
}

func (pc *PublicController) UpdatePublic(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: UpdatePublic API called")

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf(constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	var public structures.Public
	err = json.NewDecoder(r.Body).Decode(&public)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDecodingRequestBody, err)
		http.Error(w, constants.ErrInvalidInput, http.StatusBadRequest)
		return
	}
	public.ID = id

	if err := pc.repo.UpdatePublic(&public); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorUpdatingPublic, err)
		http.Error(w, constants.ErrUpdatingPublic, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(public)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogPublicCreateSuccessfully, id)
}

func (pc *PublicController) DeletePublic(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: DeletePublic API called")

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	if err := pc.repo.DeletePublic(id); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDeletingPublic, err)
		http.Error(w, constants.ErrDeletePublic, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("CONTROLLER: Public delete successfully: %v", id)
}
