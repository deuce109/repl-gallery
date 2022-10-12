package database

import (
	"database/sql"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteConnection struct {
	db *sql.DB
}

func (conn *SqliteConnection) SelectSQL(query string, params ...interface{}) ([]interface{}, error) {
	return nil, nil
}

func (conn *SqliteConnection) Get(delegate func(interface{}) bool) ([]interface{}, error) {
	return nil, nil
}

func (conn *SqliteConnection) Create(name string, valueMap map[string]string) (ok bool, err error) {
	stmt, err := conn.db.Prepare("CREATE TABLE IF NOT EXISTS ?(?)")
	if err != nil {
		return
	}

	res, err := stmt.Exec(name, generateTableDef(valueMap))
	if err != nil {
		return
	}

	fmt.Printf("%v", res)
	return true, nil
}

func InitSqlite(connectionString string) (*SqliteConnection, error) {
	db, err := sql.Open("sqlite3", connectionString)
	return &SqliteConnection{db}, err
}

func generateTableDef(valueMap map[string]string) (query string) {
	for key, value := range valueMap {
		query += fmt.Sprintf("%s %s", key, value)
	}
	return
}
