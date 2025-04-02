package metadata

import (
	"database/sql"
	"fmt"
)

func PingDB() {
	db, err := sql.Open("sqlite3", "metadata.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `SELECT * FROM video_metadata;`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title, description string
		var duration int
		if err := rows.Scan(&id, &title, &description, &duration); err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "title:", title, "description:", description, "duration:", duration)
	}
}
