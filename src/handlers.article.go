package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	//call context to render a template
	c.HTML(
		//set status 200 to OK
		http.StatusOK,
		//use the index.html template
		"index.html",
		//supply the title and payload
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
