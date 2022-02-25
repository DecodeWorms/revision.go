package storage

import (
	"database/sql"
	"fmt"
	"reflect"

	"revision.go/types"
)

type Students struct {
	d *Conn
}

func NewStudents(c *Conn) Students {
	return Students{
		d: c,
	}
}

func (s Students) Create(u types.User) (sql.Result, error) {
	var res *types.User
	var err error
	res, err = SaveData(u)

	var row sql.Result

	row, err = s.d.Client.Exec("INSERT INTO users VALUES($1,$2,$3)", res.Id, res.Name, res.Gender)
	if err != nil {
		panic(err)
		//return nil, err
	}

	return row, nil

}

func (s Students) GetUser(n string) (*types.User, error) {
	var name string
	var err error
	name, err = CheckNameLength(n)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	var res *sql.Row
	res = s.d.Client.QueryRow("SELECT id,name,gender FROM users WHERE name = $1", name)
	var us types.User

	if err = res.Scan(&us.Id, &us.Name, &us.Gender); err != nil {
		return nil, err
	}
	return &us, nil

}

func (s Students) GetUsers() ([]types.User, error) {
	var err error
	var user []types.User

	var rows *sql.Rows
	rows, err = s.d.Client.Query("SELECT id,name,gender FROM users")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var u types.User

		if err = rows.Scan(&u.Id, &u.Name, &u.Gender); err != nil {
			return nil, err
		}
		user = append(user, u)
	}
	fmt.Println(user)
	var val []types.User
	val, err = GetSlice(user)

	if err != nil {
		panic(err)
	}

	return val, nil
}

func (s Students) Update(old, new string) (sql.Result, error) {
	var err error
	var n, o string
	o, n, err = CheckForTwoStrings(new, old)
	var row sql.Result
	row, err = s.d.Client.Exec("UPDATE users SET name = $1 WHERE name = $2", o, n)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (s Students) Delete(n string) (sql.Result, error) {
	var err error
	var name string
	name, err = CheckNameLength(n)
	if err != nil {
		panic(err)
	}

	var row sql.Result
	row, err = s.d.Client.Exec("DELETE FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func CheckNameLength(n string) (string, error) {
	var err error
	if len(n) < 1 {
		return "", err
	}
	return n, nil

}

func CheckDatatype(n string) (string, error) {
	var err error
	if reflect.TypeOf(n).Kind() == reflect.String {
		return n, nil
	}
	return "", err

}

func SaveData(u types.User) (*types.User, error) {
	var res types.User
	var err error

	if u.Id != 0 && u.Name != "" && u.Gender != "" {
		res = types.User{
			Id:     u.Id,
			Name:   u.Name,
			Gender: u.Gender,
		}
		return &res, nil
	}

	return nil, err

}

func GetSlice(u []types.User) ([]types.User, error) {
	var err error

	for i, _ := range u {
		if u[i].Id != 0 && u[i].Name != "" && u[i].Gender != "" {
			return u, nil
		}
	}

	return nil, err

}

func CheckForTwoStrings(old, new string) (string, string, error) {
	var err error
	if old != "" && new != "" {
		return old, new, err
	}
	return "", "", nil

}
