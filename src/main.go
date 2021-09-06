package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}

//render HTML, JSON, or XML based on the Accept header of the request
//If the header does not specify, render HTML by default,
//provided that the template name is not present
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
