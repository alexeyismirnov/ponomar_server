package main

import (
	"embed"
	"strconv"

	"alexeyismirnov/ponomar_server/api"

	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var assets embed.FS

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	r.GET("/saints/:day/:month/:year", func(c *gin.Context) {
		day, err1 := strconv.Atoi(c.Param("day"))
		month, err2 := strconv.Atoi(c.Param("month"))
		year, err3 := strconv.Atoi(c.Param("year"))

		if err1 != nil || err2 != nil || err3 != nil {
			c.AbortWithStatus(500)
			return
		}

		res, err := api.GetSaints(day, month, year)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, res)
	})

	r.GET("/icons/:day/:month/:year", func(c *gin.Context) {
		day, err1 := strconv.Atoi(c.Param("day"))
		month, err2 := strconv.Atoi(c.Param("month"))
		year, err3 := strconv.Atoi(c.Param("year"))

		if err1 != nil || err2 != nil || err3 != nil {
			c.AbortWithStatus(500)
			return
		}

		res, err := api.GetIcons(day, month, year)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, res)
	})

	r.POST("/pericope", func(c *gin.Context) {
		var params api.Pericope
		c.BindJSON(&params)

		res, err := api.GetPericope(&params)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, res)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run()

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
