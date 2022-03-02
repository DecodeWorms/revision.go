package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"revision.go/config"
)

type Conn struct {
	Client *sql.DB
}

func NewConn(c config.Config, db *sql.DB) *Conn {
	log.Println("Connecting to DB....")
	var err error

	uri := fmt.Sprintf("host=%s dbname=%s user=%s port=%s sslmode=disable", c.DatabaseHost, c.DatabaseName, c.DatabaseUserName, c.DatabasePort)

	db, err = sql.Open("postgres", uri)

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to connect to DB %V", err))
		fmt.Println(e)
	}

	log.Println("Connected to DB..")

	return &Conn{
		Client: db,
	}
}

func Comm() {

}
