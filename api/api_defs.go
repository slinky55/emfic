package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/slinky55/emfic/db"
	"github.com/slinky55/emfic/models"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ClientCreate(c *gin.Context) {
	var data models.Client

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	client, err := models.CreateClient(data.Name, data.PayRate)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res := db.Connection.Create(&client)

	if res.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}

func ClientAddHours(c *gin.Context) {
	id := c.Param("id")

	var client models.Client

	res := db.Connection.First(&client, id)

	if res.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	timeSeconds, err := strconv.Atoi(c.Param("timeSeconds"))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	client.TimeSeconds += uint(timeSeconds)

	db.Connection.Save(&client)

	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}

func ClientAddInvoice(c *gin.Context) {
	id := c.Param("id")

	var client models.Client

	res := db.Connection.First(&client, id)

	if res.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(res.Error)
		return
	}

	invoice, err := models.CreateInvoice(&client)

	res = db.Connection.Create(&invoice)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(res.Error)
		return
	}

	err = db.Connection.Model(&client).Association("Invoices").Append(&invoice)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	client.TimeSeconds = 0
	db.Connection.Save(&client)

	c.JSON(http.StatusOK, gin.H{
		"invoice": invoice,
	})
}
