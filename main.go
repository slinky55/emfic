package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/slinky55/emfic/api"
	"github.com/slinky55/emfic/db"
	"github.com/slinky55/emfic/models"
)

func main() {
	if (len(os.Args) > 1) && (os.Args[1] == "--release") {
		gin.SetMode(gin.ReleaseMode)
	}

	dataPath := "data"

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		err := os.Mkdir(dataPath, 0755)
		if err != nil {
			log.Fatalln("Failed to create data directory")
		}
	}

	log.Println("Opening database...")

	db.Connect()

	err := db.Connection.AutoMigrate(&models.Invoice{}, &models.Client{})
	if err != nil {
		log.Fatalln("Failed to migrate data")
		return
	}

	r := gin.Default()

	gApi := r.Group("/api")
	{
		gApi.GET("/ping", api.Ping)

		gApi.POST("/client/create", api.ClientCreate)
		gApi.PATCH("/client/:id/add/hours/:timeSeconds", api.ClientAddHours)
		gApi.PATCH("/client/:id/add/invoice", api.ClientAddInvoice)
	}

	err = r.Run(":7100")
	if err != nil {
		log.Fatalln(err)
	}
}
