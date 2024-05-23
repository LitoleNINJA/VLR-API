package server

import (
	scrapper "VLR-API/scrapper"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func setAPIEndpoints(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to VLR API",
		})

	})
	router.GET("/matches", getMatches)
	router.GET("/matches/:id", getMatch)
}

func getMatches(c *gin.Context) {
	status := c.Query("status")

	if status == "" || status == "live" {
		getLiveMatches(c)
	} else if status == "completed" {
		getCompletedMatches(c)
	}
}

func getLiveMatches(c *gin.Context) {
	var matches []scrapper.Match
	val, err := rdb.Get(c, "matches").Result()
	if err == nil {
		fmt.Println("Fetching from Redis")
		matches = scrapper.UnmarshalMatches(val)
	} else {
		fmt.Println("Fetching from VLR")
		matches = scrapper.GetMatchesFromVLR()
		err := rdb.Set(c, "matches", scrapper.MarshalMatches(matches), 100*time.Second).Err()
		if err != nil {
			fmt.Println("Error setting matches in redis", err)
		}
	}

	res := gin.H{
		"matches": matches,
		"count":   len(matches),
	}
	c.IndentedJSON(200, res)
}

func getCompletedMatches(c *gin.Context) {
	var matches []scrapper.Match
	val, err := rdb.Get(c, "results").Result()
	if err == nil {
		fmt.Println("Fetching from Redis")
		matches = scrapper.UnmarshalMatches(val)
	} else {
		fmt.Println("Fetching from VLR")
		matches = scrapper.GetResultsFromVLR()
		err := rdb.Set(c, "results", scrapper.MarshalMatches(matches), 100*time.Second).Err()
		if err != nil {
			fmt.Println("Error setting results in redis", err)
		}
	}

	res := gin.H{
		"matches": matches,
		"count":   len(matches),
	}
	c.IndentedJSON(200, res)

}

func getMatch(c *gin.Context) {
	c.IndentedJSON(404, gin.H{"message": "Match not found"})
}
