package server

import (
	"crypto/internal/arbitrage"
	"crypto/internal/gate"
	"crypto/internal/stex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/gate", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &gate.PriceInfo.Prices)
	})

	router.GET("/stex", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &stex.PriceInfo.Prices)
	})

	router.GET("/arbitrage", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &arbitrage.ArbitrageResponseList)
	})

	return router
}
