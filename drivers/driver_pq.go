package drivers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func Connect(host, port, user, password, dbname string) *PostgresDB {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	Postgres.SQL = db
	return Postgres
}
