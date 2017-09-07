package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/spm/model"
)

//CreateTeam (POST)
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team model.Team

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&team); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Team Resource Exist ?
	if model.DoesTeamResourceExist(&team) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateTeam(&team); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, team)
}

//GetTeams  (GET)
func GetTeams(w http.ResponseWriter, r *http.Request) {

	teams, err := model.GetTeams()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, teams)
}

//GetTeam (GET)
func GetTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Team ID")
		return
	}

	team := model.Team{ID: id}
	if err := model.GetTeam(&team); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Team not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, team)
}

//UpdateTeam (PUT)
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Team ID")
		return
	}

	var team model.Team

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&team); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	team.ID = id

	// Does Team Resource Exist ?
	if model.DoesTeamIDExist(team.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Team ID does not exist")
		return
	}
	// Does Team Name exists for another ID
	if model.DoesTeamNameExistForAnotherID(team.Name, team.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Team Name Exists for another Team ID")
		return
	}
	if err := model.UpdateTeam(&team); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, team)
}

//DeleteTeam (DELETE)
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Team ID")
		return
	}
	team := model.Team{ID: id}

	if err := model.DeleteTeam(&team); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
