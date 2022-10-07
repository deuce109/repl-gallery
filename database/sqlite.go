package database

import (
	"database/sql"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
  "fmt"
)

type SqliteConnection struct {
  db *sql.DB
}

func (conn *SqliteConnection) SelectSQL(query string, params ...interface{}) ([]interface{}, error) {
  return nil, nil
}

func (conn *SqliteConnection) Get( delegate func(interface{}) bool) ([]interface{}, error) {
  return nil, nil
}

func (conn *SqliteConnection) Create(name string, obj interface{}) (ok bool, err error) {
  stmt, err := conn.db.Prepare("CREATE TABLE IF NOT EXISTS ?(?)")
  if err != nil {
    return
  }
  res, err stmt.Exec(name, generateTableDef(obj))
  if err != nil {
    return
  }
  
} 

func Init(connectionString string) (*SqliteConnection, error) {
  db, err := sql.Open("sqlite3", connectionString)
  return &SqliteConnection{db}, err
}

func generateTableDef(obj interface{}) (string) {
  typeElem := reflect.TypeOf(obj).Elem()
  valElem := reflect.ValueOf(obj).Elem()
  tableDefs := make(map[string]string, 0)
  for i := 0; i < typeElem.NumField(); i++ {
    field := typeElem.FieldByIndex([]int{i})
    name := field.Name
    val := valElem.FieldByName(name)
    switch val.Kind() {
      case reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Bool, reflect.Uint, reflect.Uint32, reflect.Uint16, reflect.Uint8:
        tableDefs[name] = "INT"
      case  reflect.Int64:
        tableDefs[name] = "REAL"
      case reflect.String:
        tableDefs[name] = "TEXT"
      case reflect.Float32, reflect.Float64, reflect.Complex128, reflect.Complex64:
        tableDefs[name] = "NUMERIC"
      }
  }
  query := ""
  for key, value := range tableDefs {
    query += fmt.Sprintf("'%s' %s,", key, value) 
  }
  return query
}

