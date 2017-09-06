package model

import (
	"database/sql"
)

//User  (TYPE)
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Team      Team   `json:"team"`
}

//Users (TYPE)
type Users struct {
	Users []*User `json:"users"`
}

//DoesUserResourceExist (POST)
func DoesUserResourceExist(user *User) bool {

	err := db.QueryRow("SELECT id, first_name, last_name FROM spm_users WHERE email=$1",
		user.Email).Scan(&user.ID, &user.FirstName, &user.LastName)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesUserIDExist (POST)
func DoesUserIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM spm_users WHERE id=$1", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesUserEmailExistForAnotherID (PUT)
func DoesUserEmailExistForAnotherID(email string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM spm_users WHERE email=$1", email).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateUser (POST)
func CreateUser(user *User) error {

	err := db.QueryRow(
		"INSERT INTO spm_users(first_name, last_name, email, team_id) VALUES($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Email, user.Team.ID).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

//GetUsers (GET)
func GetUsers() ([]User, error) {
	rows, err := db.Query("SELECT spm_users.id, first_name, last_name, email, team_id, name FROM spm_users " +
		"inner join spm_teams on spm_teams.id = team_id")

	if err != nil {
		return nil, err
	}

	users := []User{}

	for rows.Next() {
		defer rows.Close()

		var u User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Team.ID, &u.Team.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

//GetUser (GET)
func GetUser(user *User) error {
	return db.QueryRow("SELECT first_name, last_name, email, team_id, name FROM spm_users "+
		"inner join spm_teams on spm_teams.id = team_id WHERE spm_users.id=$1",
		user.ID).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Team.ID, &user.Team.Name)
}

//GetUserByEmail (GET)
func GetUserByEmail(user *User) error {
	return db.QueryRow("SELECT spm_users.id, first_name, last_name, team_id, name FROM spm_users "+
		"inner join spm_teams on spm_teams.id = team_id WHERE email=$1",
		user.Email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Team.ID, &user.Team.Name)
}

//UpdateUser (PUT)
func UpdateUser(user *User) error {
	_, err :=
		db.Exec("UPDATE spm_users SET first_name=$1, last_name=$2, email=$3, team_id=$4 WHERE id=$5",
			user.FirstName, user.LastName, user.Email, user.Team.ID, user.ID)

	return err
}

//DeleteUser (DELETE)
func DeleteUser(user *User) error {
	_, err := db.Exec("DELETE FROM spm_users WHERE id=$1", user.ID)

	return err
}