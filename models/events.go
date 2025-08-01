package models

import (
	"time"

	"example.com/go_rest_api_backend_project/db"
)

type Event struct {
	ID          int64
	Name        string `Binding:"required"`
	Description string `Binding:"required"`
	Location    string `Binding:"required"`
	DateTime    time.Time `Binding:"required"`
	UserID      int64
}

func (e *Event) Save() error{
	query := `INSERT INTO events (name,description,location,dateTime,user_id) 
							VALUES(?,?,?,?,?)`
	stmt,err := db.DB.Prepare(query) // we can directly use exec but using prep gives better performance in certain places
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	if err != nil {
		return err
	}
	id,err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event,error) {
	query := `SELECT * FROM events` // as this is very simple query i will not use prepare like previous query
	rows, err := db.DB.Query(query) // I have used Query() coz if we want to get bunch of rows back Query is used but if you want to insert something or edit ir delete we use Exec()
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	var events []Event
	for rows.Next()	{
		var event Event
		err := rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
		if err != nil {
			return nil,err
		}
		events = append(events, event)
	}
	// rows.Next() gives you a bool value so it will keep the loop running until we have value in loop
	return events,nil
}
func GetEventById(id int64)(*Event, error){
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query,id) // QueryRow gives us a single row so we do not need to loop and get the data
	var event Event
	err := row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
	if err != nil {
		return nil,err // we return pointer to event so we can use nil otherwise it will throw error as we would need to return empty event in this case Event{}
	}
	return &event,nil
}

func (event Event) Update () error {
	query := `
	UPDATE events
	SET name = ?,description = ?,location = ?,dateTime = ?
	WHERE id = ?
	`
	stmt ,err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name,event.Description,event.Location,event.DateTime,event.ID)
	return err
}
func (event Event) Delete () error {
	query := `DELETE FROM events WHERE id = ?`
	stmt ,err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register (userId int64) error {
	query := "INSERT INTO registrations(event_id , user_id) VALUES (? , ?)"
	stmt,err := db.DB.Prepare(query) 
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID,userId)
	return err
}
func (e Event) CancelRegister (userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt,err := db.DB.Prepare(query) 
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID,userId)
	return err
}