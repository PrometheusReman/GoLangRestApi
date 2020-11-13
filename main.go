package main

import (

	"log"

	"net/http"
	
	"github.com/gin-gonic/gin"

	centralRoute "GoLangApi/app/centralRoute"
)

func main() {
	server := gin.Default()

	server.Use(Cors())

	centralRoute.GetAllRoutes(server)
	
	log.Fatal(http.ListenAndServe(":9786", server))
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		c.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Accept, Origin ,Authorization, Content-Type, X-Csrf-Token, X-Requested-With, Range")
		c.Next()
	}
}