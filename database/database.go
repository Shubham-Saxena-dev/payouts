package database

/*
This class is responsible for creating and returning the mysql database connection
*/

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

const (
	ConnectionString = "root:admin@/myDb"
	Type             = "mysql"
)

/*
Create connection to database and return db connection
*/
func newDatabaseConnection() *sql.DB {
	db, err := sql.Open(Type, ConnectionString)

	if err != nil {
		log.Fatal("Connection to database failed:", err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Unable to ping database:", err)
	}
	db.Query("DROP table Payout")
	_, err = db.Query("CREATE TABLE Payout(SellerReference int,Currency varchar(20),Amount varchar(255));")

	if err != nil {
		log.Fatal("Failed to create table: ", err)
		panic(err)
	}

	log.Info("Connected to database")

	return db
}

//Get single instance of mySql Db connection
func GetDbConnection() *sql.DB {
	if dbInstance == nil {
		once.Do(
			func() {
				dbInstance = newDatabaseConnection()
			})
	}
	return dbInstance
}
