package main

import (
	"stock-price-api/config"
	"stock-price-api/markets"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	marketClient := markets.NewAlpacaClient()

	r := gin.Default()
	r.GET("/live/:symbol", func(c *gin.Context) {
		symbol := c.Param("symbol")

		if symbol == "" {
			c.JSON(400, gin.H{
				"message": "Missing symbol in url path",
			})
			return
		}

		symbolInfo, err := marketClient.GetLatest(symbol)

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error getting symbol info, please try again later",
			})
			return
		}

		c.IndentedJSON(200, symbolInfo)
	})

	r.Run() // listen and serve on
}
