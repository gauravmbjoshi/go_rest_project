package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" //we add a _ infront of it caz we dont want it to be removed as we are not directly using the package we will be using database/sql inbuilt package from go
)

var DB *sql.DB

func InitDB(){
	var err error
	DB,err = sql.Open("sqlite3","sqlite3.db") // 1st parameter sql package we installed and 2nd parameter is source file if file does not exist it will be created to the path amd the filename specified
	if err != nil || DB == nil {
        panic(fmt.Sprintf("Failed to open database: %v", err))
    }

	DB.SetMaxOpenConns(10) // how many open connections we can have at a time with the db
	DB.SetMaxIdleConns(5) // how many connections you want which should be open while no one is using
	createTables()
}

func createTables(){
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER 
		)
	`
	// readability its best to use `` backticks` is explaining the usage of
	// backticks (`) in Go for creating multi-line strings.

	_, err := DB.Exec(createEventsTable)
	// .Exec is the function which executes the query to the database connected and gives result and err as outputs which should be handled
	if err != nil {
		panic(fmt.Sprintf("Could not create events table: %v", err))
	}
}