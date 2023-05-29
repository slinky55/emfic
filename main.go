package main

import (
	"github.com/slinky55/emfic/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"

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

	r := router.Make()

	err = r.Run(":7100")
	if err != nil {
		log.Fatalln(err)
	}
}
