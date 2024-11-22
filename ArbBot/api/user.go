package api

import "arbbot/structures"

func AddUserAPI(user structures.User) error {
	return PostToAPI("http://localhost:8080/api/users", user)
}
