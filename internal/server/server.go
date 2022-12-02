package server

import (
	"crypto/internal/arbitrage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/prices", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// c.JSON(http.StatusOK, &gate.PriceInfo.Prices)
		c.JSON(http.StatusOK, &prcAgg)
	})

	router.GET("/balance", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &blncAgg)
	})

	router.GET("/arbitrage", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &arbitrage.ArbitrageResponseList)
	})

	return router
}
