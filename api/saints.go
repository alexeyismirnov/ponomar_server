package api

import (
	"fmt"
	"time"

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

	// https://pkg.go.dev/time#Time.Equal

	d1 := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", d1.Local())
	fmt.Printf("%t\n", IsLeapYear(year))

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
