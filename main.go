package main

import (
  "fmt"
  "embed"
  "log"
  "github.com/gin-gonic/gin"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "encoding/json"
)

//go:embed assets/saints/*
var assets embed.FS

type Saint struct {
	Day       int `gorm:"day"`
  Typikon   int `gorm:"typikon"`
	Name      string `gorm:"type:string"`
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


  b, err := json.Marshal(s)
  fmt.Println(string(b))


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": string(b),
		})
	})
	r.Run()


}
