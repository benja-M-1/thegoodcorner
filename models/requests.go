package models

import (
	"database/sql"
	"log"
)

type Request struct {
	Id   int64
	Int1 int
	Int2 int
	Str1 string
	Str2 string
}

func (r *Request) Insert (db *sql.DB) {
	query := "INSERT INTO requests(int1, int2, str1, str2) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, r.Int1, r.Int2, r.Str1, r.Str2).Scan(&r.Id)
	if err != nil {
		log.Fatal(err)
	}
}
