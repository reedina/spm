package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/spm/ctrl"
	"github.com/reedina/spm/model"

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
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.UpdateTeam).Methods("PUT")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.DeleteTeam).Methods("DELETE")

	//model.User struct
	a.Router.HandleFunc("/api/user", ctrl.CreateUser).Methods("POST")
	a.Router.HandleFunc("/api/users", ctrl.GetUsers).Methods("GET")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.GetUser).Methods("GET")
	a.Router.HandleFunc("/api/user/{email}", ctrl.GetUserByEmail).Methods("GET")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.DeleteUser).Methods("DELETE")

}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}