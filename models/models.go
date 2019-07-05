package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Datastore interface {
	CreateRequest(r *FizzbuzzRequest) (*FizzbuzzRequest, error)
	AllStatistics() ([]*Statistic, error)
}

type DB struct {
	*sql.DB
}

// Creates and return a new DB connection
func NewDB(username string, password string, database string, hostname string, port string) (*DB, error) {
	var err error

	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, hostname, port, database))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
