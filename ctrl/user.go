package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/spm/model"
)

//CreateUser (POST)
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Account Attribute Resource Exist ?
	if model.DoesUserResourceExist(&user) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	// Does Team Resource Exist ?
	if model.DoesTeamIDExist(user.Team.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Team ID does not exist")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Team name for Team id
	team := model.Team{}
	team.ID = user.Team.ID
	model.GetTeam(&team)
	user.Team.Name = team.Name

	respondWithJSON(w, http.StatusCreated, user)
}

//GetUsers  (GET)
func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := model.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

//GetUser (GET)
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	user := model.User{ID: id}
	if err := model.GetUser(&user); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

//GetUserByEmail (GET)
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, err := mail.ParseAddress(vars["email"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Email Address")
		return
	}
	user := model.User{}
	user.Email = email.Address

	if err := model.GetUserByEmail(&user); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

//UpdateUser (PUT)
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	var user model.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	user.ID = id

	// Does Team Resource Exist ?
	if model.DoesTeamIDExist(user.Team.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Team ID does not exist")
		return
	}
	// Does User ID exist ?
	if model.DoesUserIDExist(user.ID) != true {
		respondWithError(w, http.StatusBadRequest, "User ID does not exist")
		return
	}

	// Does Email exists for another User ID
	if model.DoesUserEmailExistForAnotherID(user.Email, user.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Email Exists for another User ID")
		return
	}
	if err := model.UpdateUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Team name for Team id
	team := model.Team{}
	team.ID = user.Team.ID
	model.GetTeam(&team)
	user.Team.Name = team.Name

	respondWithJSON(w, http.StatusOK, user)
}

//DeleteUser (DELETE)
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	user := model.User{ID: id}

	if err := model.DeleteUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}