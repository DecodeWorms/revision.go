package storage

import (
	"database/sql"
	"fmt"
	"reflect"

	"revision.go/types"
)

// I advise renaming this to something like "StudentService"
type Students struct {
	d *Conn
}

//  Above comment should force this to renamed to "NewStudentService"
func NewStudents(c *Conn) Students {
	return Students{
		d: c,
	}
}

// 1- The receiver,namely "Students" should become "StudentService"

// 2- This "s.d.Client" is not readable in the sense that someone reding your code needs to 
// frequently scroll up to your "Students" to fugire our what "s.d" refers to. A good naming is something like: "s.conn."

// 3- If you have to change your tab;e's name, you would have to search and make the change in every query. 
// You should define a constant variable for table.

// 4- PANICing inside a method/function that is called by other services is a terbile idea. Don't do that! Retrun the error and let the caller handle it.

// 5- Consider renaming "types" folder to "models".

// 6- The function "SaveData" purpose does not make sense to me. Does it validate the input (that is what it seems to be doing) or it saves data?

// 7- Handle input validation at the Handler later (always keep the Service/Store layer as clean and lean as possible).

// 8- The methos "Update" and a few other ones seem to be returning raw SQL. This is wrong; please return a struct!

// Purposes of "GetSlice" and "CheckForTwoStrings" are unclear to me.

// This "CheckNameLength" and "CheckDatatype" will never return error should the check fail. Why? because, you are not retuning defining and retuning error the right way.


func (s Students) Create(u types.User) (sql.Result, error) {
	var res *types.User
	var err error
	res, err = SaveData(u)

	var row sql.Result

	row, err = s.d.Client.Exec("INSERT INTO users VALUES($1,$2,$3)", res.Id, res.Name, res.Gender)
	if err != nil {
		panic(err) // terrible idea
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
