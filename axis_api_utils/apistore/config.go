package apistore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //...
)

//InitDatabase ...
func InitDatabase(username string, password string, host string, port string, dbname string) (*sql.DB, error) {

	dbLoginInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	conn, err := sql.Open("postgres", dbLoginInfo)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
