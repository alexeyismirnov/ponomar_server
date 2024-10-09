package api

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Icon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getIconData(gdb *gorm.DB, d time.Time) (saints []Icon, err error) {
	var icons, link_icons []Icon
	day := d.Day()
	month := int(d.Month())

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

func GetIcons(day int, month int, year int) (saints []Icon, err error) {
	const dbName = "assets/icons/icons.sqlite"

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	d := DateTime(day, month, year)
	leapStart := LeapStart(year)
	leapEnd := LeapEnd(year)

	if IsLeapYear(year) {
		if (d.Equal(leapStart) || d.After(leapStart)) && d.Before(leapEnd) {
			return getIconData(gdb, d.AddDate(0, 0, 1))
		} else if d.Equal(leapEnd) {
			return getIconData(gdb, leapStart)
		} else {
			return getIconData(gdb, d)
		}
	} else {
		s, err := getIconData(gdb, d)
		if err == nil && d.Equal(leapEnd) {
			s1, err := getIconData(gdb, LeapStart(2000))
			return append(s, s1...), err
		}

		return s, err
	}

}
