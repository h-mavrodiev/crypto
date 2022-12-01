package server

import (
	"crypto/internal/arbitrage"
	"crypto/internal/gate"
	"crypto/internal/stex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server(gatePriceInfo *gate.SafePrices, stexPriceInfo *stex.SafePrices, arbList *[]arbitrage.ArbitrageInfo) *gin.Engine {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/gate", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &gatePriceInfo.Prices)
	})

	router.GET("/stex", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &stexPriceInfo.Prices)
	})

	router.GET("/arbitrage", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, &arbList)
	})

	return router
}
