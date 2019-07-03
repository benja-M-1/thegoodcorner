package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Datastore interface {

}

func NewDB(username string, password string, database string, hostname string, port string) (*sql.DB, error) {
	var err error

	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, hostname, port, database))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
