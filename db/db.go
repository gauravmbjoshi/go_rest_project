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
	
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			PASSWORD TEXT NOT NULL
			 
		)
	`
	_, err := DB.Exec(createUsersTable)
	// .Exec is the function which executes the query to the database connected and gives result and err as outputs which should be handled
	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	// readability its best to use `` backticks` is explaining the usage of
	// backticks (`) in Go for creating multi-line strings.

	_, err = DB.Exec(createEventsTable)
	// .Exec is the function which executes the query to the database connected and gives result and err as outputs which should be handled
	if err != nil {
		panic(fmt.Sprintf("Could not create events table: %v", err))
	}

	createRegistration := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistration)
	if err != nil {
		panic(fmt.Sprintf("Could not create registrations table: %v", err))
	}

}

// Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())
// In the previous lectures, we started sending SQL commands to the SQLite database.

// And we did this by following different approaches:

// DB.Exec() (when we created the tables)

// Prepare() + stmt.Exec() (when we inserted data into the database)

// DB.Query() (when we fetched data from the database)

// Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

// The difference between those two methods then just is whether you're fetching data from the database (=> use Query()) or your manipulating the database / data in the database (=> use Exec()).

// But what's the advantage of using Prepare()?

// Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

// This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

// And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.

// But in order to show you the different ways of using the sql package, I decided to also include this preparation approach in this course.

