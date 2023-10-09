package main

import (
	"stock-price-api/config"
	"stock-price-api/markets"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()
	marketClient := markets.NewPolygonClient(config)	

	r := gin.Default()
	r.GET("/live/:symbol", func(c *gin.Context) {
		symbol := c.Param("symbol")

		if symbol == "" {
			c.JSON(400, gin.H{
				"Error": "Please provide a symbol",
			})
			return
		}

		symbolInfo, err := marketClient.GetLatest(symbol)

		if err != nil {
			c.JSON(500, gin.H{
				"Error while fetching symbol data, please try again later": err,
			})
			return
		}
	})

	r.Run() // listen and serve on
}
