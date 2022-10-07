package sqlite

import (
	"database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type SqliteConnection struct {
  db *sql.DB
}

func (conn SqliteConnection) SelectSQL(query string, params ...interface{}) ([]interface{}, error) {
  return nil, nil
}

func (conn SqliteConnection) Get( delegate func(interface{}) bool) ([]interface{}, error) {
  return nil, nil
}

func Init(connectionString string) (*SqliteConnection, error) {
  db, err := sql.Open("sqlite3", connectionString)
  return &SqliteConnection{db}, err
}

