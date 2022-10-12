package database

type dbConnection interface {
	Get(delegate func(interface{}) bool) ([]interface{}, error)
	SelectSQL(queryString string, params ...interface{}) ([]interface{}, error)
	Create(tableName string, columnDefs map[string]string) (bool, error)
}

type Database struct {
	ConnectionString string
	connection       dbConnection
}

func (db *Database) Get(delegate func(interface{}) bool) ([]interface{}, error) {
	return db.connection.Get(delegate)
}

func (db *Database) Select(queryString string, params ...interface{}) ([]interface{}, error) {
	return db.connection.SelectSQL(queryString, params...)
}

type DatabaseConfig struct {
	ConnectionString string `argparse:"-c,--connectionString"`
	Driver           string `argparse:"-d,--driver"`
}

func Connect(config *DatabaseConfig) (db *Database, err error) {
	var conn dbConnection
	switch config.Driver {
	case "sqlite":
		conn, err = InitSqlite(config.ConnectionString)
	default:
		conn, err = InitSqlite(":memory:")
	}

	return &Database{config.ConnectionString, conn}, err
}
