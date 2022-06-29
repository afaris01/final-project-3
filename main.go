package main

import (
	"final-project-3/database"
	"final-project-3/router"
	"os"
)

func main() {
	database.MulaiDB()
	r := router.MulaiApp()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
