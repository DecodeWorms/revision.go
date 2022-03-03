package handleers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"revision.go/models"
	"revision.go/storage"
	"revision.go/util"
)

type StudentsHandlers struct {
	srv storage.StudentServices
}

const success string = "successful"

func NewStudentServiceHandlers(service storage.StudentServices) StudentsHandlers {
	return StudentsHandlers{
		srv: service,
	}
}

func (s StudentsHandlers) Create(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Unable to UnMarshal JSON")
		return
	}

	err = ValidateData(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = s.srv.Create(user)
	if err != nil {
		log.Fatal(err)
		return

	}
	json.NewEncoder(w).Encode(success)

}

func (s StudentsHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var user *models.User
	var err error
	params := mux.Vars(r)
	name := params["name"]

	err = ValidateParameter(name)
	if err != nil {
		log.Fatal(err)
		return
	}

	user, err = s.srv.GetUser(name)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(success)
	json.NewEncoder(w).Encode(user)
}

func (s StudentsHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error

	var users []models.User
	users, err = s.srv.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(success)
	json.NewEncoder(w).Encode(users)
}

func (s StudentsHandlers) Update(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var user models.User
	var err error
	params := mux.Vars(r)
	name := params["name"]

	err = ValidateParameter(name)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}

	err = s.srv.Update(user.Name, name)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(success)

}

func (s StudentsHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	params := mux.Vars(r)
	name := params["name"]

	err = ValidateParameter(name)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = s.srv.Delete(name)
	if err != nil {
		log.Fatal(err)
		fmt.Print(err)
		return
	}
	json.NewEncoder(w).Encode(success)

}

// Helper functions are below

func ValidateData(u models.User) error {

	if u.Id == 0 {
		return errors.New("an empty id passed")
	}
	if u.Name == "" {
		return errors.New("an empty name passed")
	}

	if u.Gender == "" {
		return errors.New("an empty gender passed")
	}
	return nil

}

func ValidateParameter(n string) error {

	if n == "" {
		return errors.New("an empty paramter passed")
	}
	return nil

}

func Hello(n string) string {
	return n
}

func Calc(n int) (result int) {
	result = n * 2 / n
	return result

}

func FruitsPrice(f [4]int) (result int) {
	for i := 0; i < 4; i++ {
		result += f[i]

	}
	return result
}
