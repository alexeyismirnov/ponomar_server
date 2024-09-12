package main

import (
	"embed"
	"strconv"

	"alexeyismirnov/ponomar_server/api"

	"github.com/gin-gonic/gin"
)

//go:embed assets/saints/*
var assets embed.FS

func main() {

	/*
	   rows, err := gdb.Model(&Saint{}).Rows()

	   for rows.Next() {
	        gdb.ScanRows(rows, &s)
	        fmt.Println(s)
	     }

	*/

	r := gin.Default()

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

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run()

}
