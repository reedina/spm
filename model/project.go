package model

import (
	"database/sql"
)

//Project  (TYPE)
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Team Team   `json:"team"`
}

//Projects (TYPE)
type Projects struct {
	Projects []*Project `json:"projects"`
}

//DoesProjectResourceExist (POST)
func DoesProjectResourceExist(project *Project) bool {

	err := db.QueryRow("SELECT id, name FROM spm_projects WHERE name=$1",
		project.Name).Scan(&project.ID, &project.Name)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesProjectIDExist (POST)
func DoesProjectIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM spm_projects WHERE id=$1", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesProjectNameExistForAnotherID (PUT)
func DoesProjectNameExistForAnotherID(name string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM spm_projects WHERE name=$1", name).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateProject (POST)
func CreateProject(project *Project) error {

	err := db.QueryRow("INSERT INTO spm_projects (name, team_id) VALUES ($1, $2) RETURNING id", project.Name, project.Team.ID).
		Scan(&project.ID)

	if err != nil {
		return err
	}

	return nil
}

//GetProjects (GET)
func GetProjects() ([]Project, error) {

	rows, err := db.Query("SELECT spm_projects.id, spm_projects.name, spm_teams.id, spm_teams.name FROM spm_projects " +
		"inner join spm_teams on spm_teams.id = team_id")

	if err != nil {
		return nil, err
	}

	projects := []Project{}

	for rows.Next() {
		defer rows.Close()

		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Team.ID, &p.Team.Name); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

//GetProject (GET)
func GetProject(project *Project) error {
	return db.QueryRow("SELECT spm_projects.name, team_id, spm_teams.name FROM spm_projects "+
		"inner join spm_teams on spm_teams.id = team_id WHERE spm_projects.id=$1", project.ID).
		Scan(&project.Name, &project.Team.ID, &project.Team.Name)
}

//GetProjectByName (GET)
func GetProjectByName(project *Project) error {
	return db.QueryRow("SELECT spm_projects.id, spm_teams.id, spm_teams.name FROM spm_projects "+
		"inner join spm_teams on spm_teams.id = team_id WHERE spm_projects.name=$1",
		project.Name).Scan(&project.ID, &project.Team.ID, &project.Team.Name)
}

//GetProjectsByTeamName (GET)
func GetProjectsByTeamName(project *Project) ([]Project, error) {
	rows, err := db.Query("SELECT spm_projects.id, spm_projects.name, spm_teams.id, spm_teams.name FROM spm_projects "+
		"inner join spm_teams on spm_teams.id = team_id WHERE spm_teams.name=$1",
		project.Team.Name)

	if err != nil {
		return nil, err
	}

	projects := []Project{}

	for rows.Next() {
		defer rows.Close()

		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Team.ID, &p.Team.Name); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

//GetProjectsByTeamID (GET)
func GetProjectsByTeamID(project *Project) ([]Project, error) {
	rows, err := db.Query("SELECT spm_projects.id, spm_projects.name, spm_teams.id, spm_teams.name FROM spm_projects "+
		"inner join spm_teams on spm_teams.id = team_id WHERE spm_teams.id=$1",
		project.Team.ID)

	if err != nil {
		return nil, err
	}

	projects := []Project{}

	for rows.Next() {
		defer rows.Close()

		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Team.ID, &p.Team.Name); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

//UpdateProject (PUT)
func UpdateProject(project *Project) error {
	_, err :=
		db.Exec("UPDATE spm_projects SET name=$1, team_id=$2 WHERE id=$3", project.Name, project.Team.ID, project.ID)

	return err
}

//DeleteProject (DELETE)
func DeleteProject(project *Project) error {
	_, err := db.Exec("DELETE FROM spm_projects WHERE id=$1", project.ID)

	return err
}
