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
				"Error": "Error getting symbol info, please try again later",
			})
		}

		c.IndentedJSON(200, symbolInfo)
	})

	r.Run() // listen and serve on
}
