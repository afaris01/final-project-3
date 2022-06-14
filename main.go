package main

import (
	"final-project-3/database"
	"final-project-3/router"
)

func main() {
	database.MulaiDB()
	r := router.MulaiApp()
	r.Run(":8080")
}