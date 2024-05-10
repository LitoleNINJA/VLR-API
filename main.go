package main

import (
	"net/http"

	scrapper "VLR-CLI/api/scrapper"

	"github.com/gin-gonic/gin"
)

func WelcomeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Valorant API",
	})
}

func GetMatches(c *gin.Context) {
	matches := scrapper.GetMatches()
	c.JSON(http.StatusOK, matches)
}

func main() {
	router := gin.Default()
	router.GET("/", WelcomeMessage)
	router.GET("/matches", GetMatches)
	router.Run()
}
