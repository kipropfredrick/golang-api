package db

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//create struct of type database

type Database struct {
	gorm.Model
	db *sql.DB
}
// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = "kiprop"
// dbname   = "mydb"
func NewDatabase() (*Database, error) {
	db,err := sql.Open("postgres","postgresql://godbtest:31877101@go-chat.cydlpjqntwzu.us-west-2.rds.amazonaws.com:5432/postgres?sslmode=disable");
	 if err != nil {
		return nil, err
	 }
	 //return pointer database struct vhvhg
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