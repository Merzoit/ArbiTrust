package controllers

import (
	"arb/constants"
	"arb/repositories"
	"arb/structures"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userRepo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{userRepo: repo}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallCreateUser)

	var user structures.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDecodingRequestBody, err)
		http.Error(w, constants.LogErrorDecodingRequestBody, http.StatusBadRequest)
		return
	}

	if err := uc.userRepo.AddUser(&user); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorAddingUser, err)
		http.Error(w, constants.ErrAddingUser, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)

	}

	log.Printf("CONTROLLER: "+constants.LogUserCreateSuccessfully, user.ID)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallGetUserByID)

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusNotFound)
		return
	}

	user, err := uc.userRepo.GetUserById(id)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorFetchingUser, err)
		http.Error(w, constants.ErrFetchingTeams, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
		return
	}

	log.Printf("CONTROLLER: "+constants.LogUserFetchSuccessfully, user.ID)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallUpdateUser)

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	var user structures.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDecodingRequestBody, err)
		http.Error(w, constants.ErrInvalidInput, http.StatusBadRequest)
		return
	}
	user.ID = id

	if err := uc.userRepo.UpdateUser(&user); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorUpdatingUser, err)
		http.Error(w, constants.ErrUpdatingUser, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorEncodingResponse, err)
	}

	log.Printf("CONTROLLER: "+constants.LogUserUpdateSuccessfully, user.ID)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER: " + constants.CallDeleteUser)

	id, err := extractID(r.URL.Path)
	if err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorExtractingID, err)
		http.Error(w, constants.ErrInvalidID, http.StatusBadRequest)
		return
	}

	if err := uc.userRepo.DeleteUser(id); err != nil {
		log.Printf("CONTROLLER: "+constants.LogErrorDeletingUser, err)
		http.Error(w, constants.ErrDeleteUser, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("CONTROLLER: "+constants.LogUserDeleteSuccessfully, id)
}

func extractID(path string) (uint, error) {
	log.Println("CONTROLLER: " + constants.CallExtractID)

	parts := strings.Split(path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		log.Printf("CONTROLLER: " + constants.ErrFailedToExtractFromPath)
		return 0, err
	}
	log.Printf("CONTROLLER: "+constants.LogExtractIDSuccessfully, path)
	return uint(id), nil
}
