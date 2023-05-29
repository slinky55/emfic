package router

import (
	"github.com/gin-gonic/gin"
	"github.com/slinky55/emfic/api"
)

func Make() *gin.Engine {
	r := gin.Default()

	gApi := r.Group("/api")
	{
		gApi.GET("/ping", api.Ping)

		gApi.POST("/client/create", api.ClientCreate)
		gApi.PATCH("/client/:id/add/time/:timeSeconds", api.ClientAddHours)
		gApi.PATCH("/client/:id/add/invoice", api.ClientAddInvoice)
	}

	return r
}
