package api

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Icon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetIcons(day int, month int, year int) (saints []Icon, err error) {
	const dbName = "assets/icons/icons.sqlite"

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var icons []Icon

	if result := gdb.Where("day = ? AND month = ? AND has_icons = 1", day, month).Find(&icons); result.Error != nil {
		return nil, result.Error
	}

	return icons, nil

}
