package metadata

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type VideoMetadata struct {
	ID          int
	Title       string
	Description string
	Duration    int
}

// Insert mock data
func InsertMockData() {
	db, err := sql.Open("sqlite3", "metadata.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO video_metadata (title, description, duration) VALUES
		('Sample Video 1', 'This is a sample video description.', 120),
		('Sample Video 2', 'This is another sample video description.', 150);`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Mock data inserted successfully")
}

func InitDb() {
	db, err := sql.Open("sqlite3", "metadata.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS video_metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		duration INTEGER
	);`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database initialized successfully")
}
