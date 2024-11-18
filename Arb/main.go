package main

import (
	"arb/controllers"
	"arb/db"
	"arb/repositories"
	"log"
	"net/http"
)

func main() {

	db.InitDB()
	defer db.DatabasePool.Close()

	userRepo := repositories.NewPgUserRepository()
	userController := controllers.NewUserController(userRepo)

	teamRepo := repositories.NewPgTeamRepository()
	teamController := controllers.NewTeamController(teamRepo)

	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.CreateUser(w, r)
		case http.MethodGet:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userController.GetUserById(w, r)
		case http.MethodPut:
			userController.UpdateUser(w, r)
		case http.MethodDelete:
			userController.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/teams", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			teamController.CreateTeam(w, r)
		case http.MethodGet:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/teams/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			teamController.GetTeamById(w, r)
		case http.MethodPut:
			teamController.UpdateTeam(w, r)
		case http.MethodDelete:
			teamController.DeleteTeam(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/allteams/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			teamController.GetAllTeams(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("OK!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
