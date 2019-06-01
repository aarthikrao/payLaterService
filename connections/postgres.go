package connections

import (
	"github.com/aarthikrao/payLaterService/models"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/aarthikrao/payLaterService/config"
)

// More about pgDB https://godoc.org/github.com/go-pg/pg

// PGDB is the singleton instance of DB connection
var PGDB *pg.DB

// InitDatabase : start db connections
func InitDatabase() {
	var err error
	PGDB, err = newDBCOnnection("psqlInfo")
	if err != nil {
		panic(err)
	}
	createSchema(PGDB)
}

// NewDBCOnnection makes a new database using the connection string and
// returns it, otherwise returns the error
func newDBCOnnection(connString string) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     config.Config.GetString("db.user"),
		Addr:     config.Config.GetString("db.host"),
		Password: config.Config.GetString("db.password"),
		Database: config.Config.GetString("db.dbname"),
	})
	return db, nil
}

// Creates the table if it doesnt exisist
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*models.User)(nil), (*models.Merchant)(nil), (*models.Transaction)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// ShutDown : close connections dusing shutdown
func ShutDown() {
	PGDB.Close()
}
