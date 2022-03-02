package storage

import (
	"database/sql"

	"revision.go/models"
)

const tableName = "users"

type StudentServices struct {
	con *Conn
}

func NewStudents(c *Conn) StudentServices {
	return StudentServices{
		con: c,
	}
}

func (s StudentServices) Create(u models.User) error {

	var err error

	_, err = s.con.Client.Exec("INSERT INTO users VALUES($1,$2,$3) ", u.Id, u.Name, u.Gender)
	if err != nil {
		return err
	}
	return nil

}

func (s StudentServices) GetUser(n string) (*models.User, error) {
	var name string
	var err error

	var res *sql.Row
	res = s.con.Client.QueryRow("SELECT id,name,gender FROM users WHERE name = $1", name)
	var us models.User

	if err = res.Scan(&us.Id, &us.Name, &us.Gender); err != nil {
		return nil, err
	}
	return &us, nil

}

func (s StudentServices) GetUsers() ([]models.User, error) {
	var err error
	var user []models.User

	var rows *sql.Rows
	rows, err = s.con.Client.Query("SELECT id,name,gender FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u models.User

		if err = rows.Scan(&u.Id, &u.Name, &u.Gender); err != nil {
			return nil, err
		}
		user = append(user, u)
	}
	return user, nil
}

func (s StudentServices) Update(old, new string) error {
	var err error
	_, err = s.con.Client.Exec("UPDATE users SET name = $1 WHERE name = $2", old, new)
	if err != nil {
		return err
	}
	return nil
}

func (s StudentServices) Delete(n string) error {
	var err error

	//var row sql.Result
	_, err = s.con.Client.Exec("DELETE FROM users WHERE name = $1", n)
	if err != nil {
		return err
	}
	return nil
}
