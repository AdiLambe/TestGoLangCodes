/*
Gin is a powerful HTTP Web framework written in golang.
it simplifies building web application, handling routing, middleware
and request/reponse processing.
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin!"})
	})
	r.Run(":8080")
}

/*we import the "github.com/gin-gonic/gin" package to use Gin
The gin.Default() func creates a new gin router with default middleware
we define a route for the root path ("/") using r.GET(...)
The router handler responds with a JSON message "message": "Hello Gin!"
Finally, we run the server on port 8080 using "r.Run(":8080")"

Advanced usage : Gin provides powerful features like middleware
routing groups, and parameter handling.

Built-in middleware in gin:
	gin provides built-in middleware like gin.logger()
	(for request logging) and gin.Recovery()(for handling panics)
	You can use them with r.Use(gin.logger()) and r.Use(gin.Recovery())*/

/* Context contains all the information about the request that the handler might need to process it.*/
