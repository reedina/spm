package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/spm/model"
)

//CreateProject (POST)
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Project Resource Exist ?
	if model.DoesProjectResourceExist(&project) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	// Does Team Resource Exist ?
	if model.DoesTeamIDExist(project.Team.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Team ID does not exist")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateProject(&project); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Team name for Team id
	team := model.Team{}
	team.ID = project.Team.ID
	model.GetTeam(&team)
	project.Team.Name = team.Name

	respondWithJSON(w, http.StatusCreated, project)
}

//GetProjects  (GET)
func GetProjects(w http.ResponseWriter, r *http.Request) {

	projects, err := model.GetProjects()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, projects)
}

//GetProject (GET)
func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Project ID")
		return
	}

	project := model.Project{ID: id}
	if err := model.GetProject(&project); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Project not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, project)
}

//GetProjectByName (GET)
func GetProjectByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectName := vars["name"]

	project := model.Project{}
	project.Name = projectName

	if err := model.GetProjectByName(&project); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Project not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, project)
}

//UpdateProject (PUT)
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Project ID")
		return
	}

	var project model.Project

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	project.ID = id

	// Does Team Resource Exist ?
	if model.DoesTeamIDExist(project.Team.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Team ID does not exist")
		return
	}

	// Does Project ID exist ?
	if model.DoesProjectIDExist(project.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Project ID does not exist")
		return
	}

	// Does Project Name exists for another ID
	if model.DoesProjectNameExistForAnotherID(project.Name, project.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Project Name Exists for another Project ID")
		return
	}

	if err := model.UpdateProject(&project); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Team name for Team id
	team := model.Team{}
	team.ID = project.Team.ID
	model.GetTeam(&team)
	project.Team.Name = team.Name

	respondWithJSON(w, http.StatusOK, project)
}

//DeleteProject (DELETE)
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Project ID")
		return
	}
	project := model.Project{ID: id}

	if err := model.DeleteProject(&project); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
