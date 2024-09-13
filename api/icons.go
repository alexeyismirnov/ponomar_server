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

	var icons, link_icons []Icon

	if result := gdb.Where("day = ? AND month = ? AND has_icon = 1", day, month).Find(&icons); result.Error != nil {
		return nil, result.Error
	}

	if result := gdb.Model(&Icon{}).Joins(
		"JOIN link_icons ON icons.id = link_icons.id",
	).Select(
		"icons.id AS id, link_icons.name AS name",
	).Where(
		"link_icons.day = ? AND link_icons.month = ? AND icons.has_icon = 1", day, month,
	).Scan(&link_icons); result.Error != nil {
		return nil, result.Error
	}

	return append(icons, link_icons...), nil

}
