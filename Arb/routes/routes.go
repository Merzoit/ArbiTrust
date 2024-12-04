package routes

import (
	"arb/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, userController *controllers.UserController) {
	userRoutes := router.PathPrefix("/api/user").Subrouter()
	userRoutes.HandleFunc("", userController.CreateUser).Methods(http.MethodPost)
	userRoutes.HandleFunc("/{id}", userController.GetUserById).Methods(http.MethodGet)
	userRoutes.HandleFunc("/{id}", userController.UpdateUser).Methods(http.MethodPut)
	userRoutes.HandleFunc("/{id}", userController.DeleteUser).Methods(http.MethodDelete)
}

func RegisterTeamRoutes(router *mux.Router, teamController *controllers.TeamController) {
	teamRoutes := router.PathPrefix("/api/team").Subrouter()
	teamRoutes.HandleFunc("", teamController.CreateTeam).Methods(http.MethodPost)
	teamRoutes.HandleFunc("/all", teamController.GetAllTeams).Methods(http.MethodGet)
	teamRoutes.HandleFunc("/{id}", teamController.GetTeamById).Methods(http.MethodGet)
	teamRoutes.HandleFunc("/{id}", teamController.UpdateTeam).Methods(http.MethodPut)
	teamRoutes.HandleFunc("/{id}", teamController.DeleteTeam).Methods(http.MethodDelete)
}

func RegisterPublicRoutes(router *mux.Router, publicController *controllers.PublicController) {
	publicRoutes := router.PathPrefix("/api/public").Subrouter()
	publicRoutes.HandleFunc("", publicController.CreatePublic).Methods(http.MethodPost)
	publicRoutes.HandleFunc("/all", publicController.GetAllPublics).Methods(http.MethodGet)
	publicRoutes.HandleFunc("/{id}", publicController.GetPublicByID).Methods(http.MethodGet)
	publicRoutes.HandleFunc("/{id}", publicController.UpdatePublic).Methods(http.MethodPut)
	publicRoutes.HandleFunc("/{id}", publicController.DeletePublic).Methods(http.MethodDelete)
}
