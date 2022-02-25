package handleers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"revision.go/storage"
	"revision.go/types"
)

type StudentsHandlers struct {
	st storage.Students
}

func NewStudentsHandlers(st storage.Students) StudentsHandlers {
	return StudentsHandlers{
		st: st,
	}
}

func (s StudentsHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var user types.User
	var row sql.Result

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Unable to create resources")
	}

	row, err = s.st.Create(user)
	if err != nil {
		json.NewEncoder(w).Encode("can not create resources")
	}
	json.NewEncoder(w).Encode(row)
}

func (s StudentsHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *types.User

	params := mux.Vars(r)
	name := params["name"]
	user, err = s.st.GetUser(name)
	if err != nil {
		json.NewEncoder(w).Encode("Results are not found")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s StudentsHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	var users []types.User
	users, err = s.st.GetUsers()

	if err != nil {
		json.NewEncoder(w).Encode("Resources are not available")
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (s StudentsHandlers) Update(w http.ResponseWriter, r *http.Request) {
	var user types.User
	var err error
	params := mux.Vars(r)
	name := params["name"]

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Can not Unmarshal JSON")
		return
	}
	var row sql.Result
	row, err = s.st.Update(name, user.Name)
	if err != nil {
		json.NewEncoder(w).Encode("Resources are not available")
		return
	}
	json.NewEncoder(w).Encode(row)

}

func (s StudentsHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	name := params["name"]

	var row sql.Result
	row, err = s.st.Delete(name)
	if err != nil {
		json.NewEncoder(w).Encode("results are not found")
	}
	json.NewEncoder(w).Encode(row)

}
