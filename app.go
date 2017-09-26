package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/spm/ctrl"
	"github.com/reedina/spm/model"
	"github.com/rs/cors"

	//Initialize pq driver
	_ "github.com/lib/pq"
)

//App  (TYPE)
type App struct {
	Router *mux.Router
}

//InitializeApplication - Init router, db connection and restful routes
func (a *App) InitializeApplication(user, password, url, dbname string) {

	model.ConnectDB(user, password, url, dbname)
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

//InitializeRoutes - Declare all application routes
func (a *App) InitializeRoutes() {

	//model.Team struct
	a.Router.HandleFunc("/api/team", ctrl.CreateTeam).Methods("POST")
	a.Router.HandleFunc("/api/teams", ctrl.GetTeams).Methods("GET")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.GetTeam).Methods("GET")
	a.Router.HandleFunc("/api/team/{name}", ctrl.GetTeamByName).Methods("GET")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.UpdateTeam).Methods("PUT")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.DeleteTeam).Methods("DELETE")

	//model.User struct
	a.Router.HandleFunc("/api/user", ctrl.CreateUser).Methods("POST")
	a.Router.HandleFunc("/api/users", ctrl.GetUsers).Methods("GET")
	a.Router.HandleFunc("/api/users/team/name/{name}", ctrl.GetUsersByTeamName).Methods("GET")
	a.Router.HandleFunc("/api/users/team/id/{id:[0-9]+}", ctrl.GetUsersByTeamID).Methods("GET")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.GetUser).Methods("GET")
	a.Router.HandleFunc("/api/user/{email}", ctrl.GetUserByEmail).Methods("GET")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.DeleteUser).Methods("DELETE")

	//model.Project struct
	a.Router.HandleFunc("/api/project", ctrl.CreateProject).Methods("POST")
	a.Router.HandleFunc("/api/projects", ctrl.GetProjects).Methods("GET")
	a.Router.HandleFunc("/api/projects/team/name/{name}", ctrl.GetProjectsByTeamName).Methods("GET")
	a.Router.HandleFunc("/api/projects/team/id/{id:[0-9]+}", ctrl.GetProjectsByTeamID).Methods("GET")
	a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.GetProject).Methods("GET")
	a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.UpdateProject).Methods("PUT")
	a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.DeleteProject).Methods("DELETE")
}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Fatal(http.ListenAndServe(addr, c.Handler(a.Router)))
}
