package main

import (
	"challenge-2-chapter-3/database"
	"challenge-2-chapter-3/router"
)

func main() {

	database.StartDB()

	router.StartApp().Run(":8080")

}
