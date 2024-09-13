package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"yehtun.com/rest-api-crud/models"
)

var DB *sql.DB

// ConnectDB establishes a connection to MySQL
func ConnectDB() {
    var err error
    dsn := "root:yeye@tcp(127.0.0.1:3306)/postdb"
    fmt.Println("Connecting to database")
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
       
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
       
    }
     // This sets the db connection in the models package
     models.SetDB(DB)
    fmt.Println("Database connection established")
}