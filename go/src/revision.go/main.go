package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"revision.go/config"
	handleers "revision.go/handlers"
	"revision.go/storage"
)

var c *storage.Conn
var s storage.Students

func init() {
	_ = godotenv.Load()
	h := os.Getenv("HOST")
	d := os.Getenv("DATABASE_NAME")
	p := os.Getenv("PORT")
	u := os.Getenv("USER_NAME")

	cnf := config.Config{
		DatabaseHost:     h,
		DatabaseName:     d,
		DatabasePort:     p,
		DatabaseUserName: u,
	}
	var dab *sql.DB
	c = storage.NewConn(cnf, dab)
	s = storage.NewStudents(c)
}
func main() {
	std := handleers.NewStudentsHandlers(s)
	router := mux.NewRouter()
	router.HandleFunc("/s/cr", std.Create).Methods(http.MethodPost)
	router.HandleFunc("/s/us/{name}", std.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/s/sts", std.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/s/up/{name}", std.Update).Methods(http.MethodPut)
	router.HandleFunc("/s/del/{name}", std.Delete).Methods(http.MethodDelete)
	//router.HandleFunc("/st/aut", std.AutoCreate).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8004", router))

}
