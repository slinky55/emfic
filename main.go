package main

import (
		"fmt"
		"github.com/gin-gonic/gin"

		"github.com/slinky55/emfic/api"
)

func main() {
		fmt.Println("Hello, World!")
		
		r := gin.Default()
		
		g_api := r.Group("/api")
		{
				g_api.GET("/ping", api.Ping)		
		}

		r.Run(":7100")
}
