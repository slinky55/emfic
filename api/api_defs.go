package api

import (
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

	res := db.Connection.Create(&data)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

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
		return
	}

	invoice, err := models.CreateInvoice(&client, db.Connection)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	client.TimeSeconds = 0
	db.Connection.Save(&client)

	err = db.Connection.Model(&client).Association("Invoice").Append(&invoice)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"invoice": invoice,
	})
}
