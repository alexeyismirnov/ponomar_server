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

	r.POST("/feofan", func(c *gin.Context) {
		var params api.FeofanParams
		c.BindJSON(&params)

		res, err := api.GetFeofan(&params)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.String(200, res)
	})

	r.GET("/tropsaint/:day/:month/:year", func(c *gin.Context) {
		day, err1 := strconv.Atoi(c.Param("day"))
		month, err2 := strconv.Atoi(c.Param("month"))
		year, err3 := strconv.Atoi(c.Param("year"))

		if err1 != nil || err2 != nil || err3 != nil {
			c.AbortWithStatus(500)
			return
		}

		res, err := api.GetSaintTroparion(day, month, year)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, res)
	})

	r.GET("/tropfeast", func(c *gin.Context) {
		id := c.Query("id")
		res, err := api.GetFeastTroparion(id)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, res)
	})

	r.GET("/bookdata", func(c *gin.Context) {
		filename := c.Query("filename")
		res, err := api.GetBookData(filename)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, res)
	})

	r.GET("/bookcontent", func(c *gin.Context) {
		filename := c.Query("filename")
		section, err1 := strconv.Atoi(c.Query("section"))
		item, err2 := strconv.Atoi(c.Query("item"))

		if err1 != nil || err2 != nil {
			c.AbortWithStatus(500)
			return
		}

		res, err := api.GetBookContent(filename, section, item)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.String(200, res)
	})

	r.GET("/bookchapters", func(c *gin.Context) {
		bookname := c.Query("bookname")
		lang := c.Query("lang")

		count, err := api.GetBookChapters(bookname, lang)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.String(200, strconv.FormatInt(count, 10))
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
