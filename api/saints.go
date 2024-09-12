package api

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Saint struct {
	Day     int    `json:"day"`
	Typikon int    `json:"typikon"`
	Name    string `json:"name"`
}

func GetSaints(day int, month int, year int) (saints []Saint, err error) {
	fmt.Printf("%d %d %d\n", day, month, year)

	dbName := fmt.Sprintf("assets/saints/saints_%02d_ru.sqlite", month)

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var s []Saint

	if result := gdb.Where("day = ?", day).Order("typikon DESC").Find(&s); result.Error != nil {
		return nil, result.Error
	}

	return s, nil
}
