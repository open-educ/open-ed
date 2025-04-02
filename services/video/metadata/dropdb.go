package metadata

import (
	"database/sql"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DropDB() {
	db, err := sql.Open("sqlite3", "metadata.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(`DROP TABLE IF EXISTS video_metadata;`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database wiped successfully")
}
