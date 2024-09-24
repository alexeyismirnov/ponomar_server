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

func getSaintData(d time.Time) (saints []Saint, err error) {
	day := d.Day()
	month := int(d.Month())

	dbName := fmt.Sprintf("assets/saints/saints_%02d_ru.sqlite", month)

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	s := []Saint{}

	if result := gdb.Where("day = ?", day).Order("typikon DESC").Find(&s); result.Error != nil {
		return nil, result.Error
	}

	return s, nil
}

func GetSaints(day int, month int, year int) (saints []Saint, err error) {
	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	leapStart := time.Date(year, time.Month(2), 29, 0, 0, 0, 0, time.UTC)
	leapEnd := time.Date(year, time.Month(3), 13, 0, 0, 0, 0, time.UTC)

	if IsLeapYear(year) {
		if (d.Equal(leapStart) || d.After(leapStart)) && d.Before(leapEnd) {
			return getSaintData(d.AddDate(0, 0, 1))
		} else if d.Equal(leapEnd) {
			return getSaintData(time.Date(year, time.Month(2), 29, 0, 0, 0, 0, time.UTC))
		} else {
			return getSaintData(d)
		}
	} else {
		s, err := getSaintData(d)
		if err == nil && d.Equal(leapEnd) {
			s1, err := getSaintData(time.Date(2000, time.Month(2), 29, 0, 0, 0, 0, time.UTC))
			return append(s, s1...), err
		}

		return s, err
	}

}
