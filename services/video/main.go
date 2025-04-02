package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	dbHandler "github.com/open-educ/open-ed/services/video/metadata"
	staticHandler "github.com/open-educ/open-ed/services/video/static"
)

func main() {
	router := gin.Default()

	router.GET("/video/:filepath", staticHandler.ServeVideo)

	if len(os.Args) > 2 {
		fmt.Println("Invalid command: use -init to initialize the database or -drop to drop the database.")
		fmt.Println(len(os.Args))
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-init":
			dbHandler.InitDb()
		case "-populate":
			dbHandler.InsertMockData()
		case "-drop":
			dbHandler.DropDB()
		case "-ping":
			dbHandler.PingDB()
		case "-help":
			fmt.Println("Commands:")
			fmt.Println("-init: Initialize the database")
			fmt.Println("-drop: Drop the database")
			fmt.Println("-ping: Ping the database")
			fmt.Println("-help: Show this help message")
		default:
			fmt.Println("Invalid command")
		}
	} else {
		router.Run(":8080")
		fmt.Println("Server started on :8080")
		return
	}
}
