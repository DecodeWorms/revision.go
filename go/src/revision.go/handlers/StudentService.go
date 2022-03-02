package handleers

import (
	"encoding/json"
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

	_, err = ValidateData(user)
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
	var err error
	var user *models.User

	params := mux.Vars(r)
	name := params["name"]

	_, err = ValidateParameter(name)
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

	_, err = ValidateUsersResults(users)

	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(success)
	json.NewEncoder(w).Encode(users)
}

func (s StudentsHandlers) Update(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	params := mux.Vars(r)
	name := params["name"]

	_, err = ValidateParameter(name)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	err = s.srv.Update(name, user.Name)
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(success)

}

func (s StudentsHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	name := params["name"]

	_, err = ValidateParameter(name)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = s.srv.Delete(name)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Errorf("Err res : %v", err))
	}
	json.NewEncoder(w).Encode(success)

}

func ValidateData(u models.User) (*models.User, error) {
	var res models.User

	if u.Id != 0 && u.Name != "" && u.Gender != "" {
		res = models.User{
			Id:     u.Id,
			Name:   u.Name,
			Gender: u.Gender,
		}
		return &res, nil

	}

	return nil, fmt.Errorf("One or all the field(is) is empty")

}

func ValidateParameter(n string) (string, error) {

	if n != "" {
		return n, nil
	}
	return "", fmt.Errorf("paramter is empty")

}

func ValidateUsersResults(u []models.User) ([]models.User, error) {

	for i, _ := range u {
		if u[i].Id != 0 && u[i].Name != "" && u[i].Gender != "" {
			return u, nil
		}
	}

	return nil, fmt.Errorf("One or all of the results are empty")

}
