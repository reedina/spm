package model

import "database/sql"

//Team  (TYPE)
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Teams (TYPE)
type Teams struct {
	Teams []*Team `json:"teams"`
}

//DoesTeamResourceExist (POST)
func DoesTeamResourceExist(team *Team) bool {

	err := db.QueryRow("SELECT id, name FROM spm_teams WHERE name=$1", team.Name).Scan(&team.ID, &team.Name)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesTeamIDExist (POST)
func DoesTeamIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM spm_teams WHERE id=$1", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//CreateTeam (POST)
func CreateTeam(team *Team) error {

	err := db.QueryRow("INSERT INTO spm_teams(name) VALUES($1) RETURNING id", team.Name).Scan(&team.ID)

	if err != nil {
		return err
	}

	return nil
}

//GetTeams (GET)
func GetTeams() ([]Team, error) {
	rows, err := db.Query("SELECT id, name FROM spm_teams")

	if err != nil {
		return nil, err
	}

	teams := []Team{}

	for rows.Next() {
		defer rows.Close()

		var t Team
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}

	return teams, nil
}

//GetTeam (GET)
func GetTeam(team *Team) error {
	return db.QueryRow("SELECT name FROM spm_teams WHERE id=$1", team.ID).Scan(&team.Name)
}

//UpdateTeam (PUT)
func UpdateTeam(team *Team) error {
	_, err :=
		db.Exec("UPDATE spm_teams SET name=$1 WHERE id=$2", team.Name, team.ID)

	return err
}

//DeleteTeam (DELETE)
func DeleteTeam(team *Team) error {
	_, err := db.Exec("DELETE FROM spm_teams WHERE id=$1", team.ID)

	return err
}