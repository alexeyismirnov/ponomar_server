package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed assets/saints/*
var assets embed.FS

type Saint struct {
	Day     int    `json:"day"`
	Typikon int    `json:"typikon"`
	Name    string `json:"name"`
}

func main() {
	dbName := "assets/saints/saints_09_ru.sqlite"

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var s []Saint
	/*
	   rows, err := gdb.Model(&Saint{}).Rows()

	   for rows.Next() {
	        gdb.ScanRows(rows, &s)
	        fmt.Println(s)
	     }

	*/

	if result := gdb.Where("day = ?", 2).Find(&s); result.Error != nil {
		log.Fatal(result.Error)
	}

	/*
		b, err := json.Marshal(s)
		fmt.Println(string(b))
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

		c.String(200, fmt.Sprintf("%d %d %d", day, month, year))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, s)
	})
	r.Run()

}
