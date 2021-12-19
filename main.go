package main

import (
	"cleverit/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api": "Test MS Beer",
		})
	})
	router.GET("/beers", controller.SearchBeers)
	router.POST("/beers", controller.AddBeers)
	router.GET("/beers/:beerID", controller.SearchBeerById)
	router.GET("/beers/:beerID/boxprice", controller.BoxBeerPriceById)

	router.Run(":8080")
}
