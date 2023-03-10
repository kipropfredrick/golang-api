package db

import (
	"database/sql"
	_"github.com/lib/pq"
)

//create struct of type database

type Database struct {
	db *sql.DB
}
// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = "kiprop"
// dbname   = "mydb"
func NewDatabase() (*Database, error) {
	db,err := sql.Open("postgres","postgresql://admin:secret@localhost:5432/go-chat?sslmode=disable");
	 if err != nil {
		return nil, err
	 }
	 //return pointer database struct
	 return &Database{db: db},nil
}
//create func to close the db
func (d *Database) Close() {
	d.db.Close()
}
//get db fucn because the database connection is encapsulated
func (d *Database) GetDB() *sql.DB {
	return d.db
}