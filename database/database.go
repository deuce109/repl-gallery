package database

import (
	"main/database/sqlite"
)

type dbConnection interface {
  Get ( delegate func(interface{}) bool) ([]interface{}, error)
  SelectSQL  ( queryString string, params ...interface{}) ([]interface{}, error)
}

type Database struct {
  ConnectionString string
  connection dbConnection
}

func (db *Database) Get (delegate func(interface{}) bool) ([]interface{}, error) {
  return db.connection.Get(delegate)
}

func (db *Database) Select (queryString string, params ...interface{}) ([]interface{}, error) {
  return db.connection.SelectSQL (queryString, params...)
}

type DatabaseConfig  struct {
  ConnectionString string `argparse:"-c,--connectionString"`
}

func Connect(config *DatabaseConfig) (*Database, error) {
  conn ,err := sqlite.Init(config.ConnectionString)
  
  return &Database{config.ConnectionString, conn}, err
}