package app

import "database/sql"

type Container struct {
	DB *sql.DB
}
