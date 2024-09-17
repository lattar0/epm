package main

import (
	"epm/models"
  "epm/routes"
)

func main() {
	models.ConnectDatabase()

	r := routes.SetupRouter()
	r.Run(":8080")
}
