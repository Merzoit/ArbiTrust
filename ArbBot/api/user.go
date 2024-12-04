package api

import (
	"arbbot/constants"
	"arbbot/structures"
	"log"
)

func AddUserAPI(user structures.User) error {
	log.Printf(constants.CallAddUser, user)
	err := PostToAPI("http://localhost:8080/api/users", user)
	if err != nil {
		log.Printf(constants.LogErrorAddingUser, user)
	}
	log.Println(constants.LogUserAddingSuccessfully)
	return nil
}
