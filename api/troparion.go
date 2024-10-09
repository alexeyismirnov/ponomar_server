package api

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Troparion struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Comment string `json:"comment"`
}

func getSaintTroparionData(gdb *gorm.DB, d time.Time) (troparia []Troparion, err error) {
	day := d.Day()
	month := int(d.Month())

	s := []Troparion{}

	if result := gdb.Table("tropari").Where("day = ? AND month = ?", day, month).Find(&s); result.Error != nil {
		return nil, result.Error
	}

	return s, nil
}

func GetSaintTroparion(day int, month int, year int) (trop []Troparion, err error) {
	dbName := "assets/troparia/troparion.sqlite"
	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	d := DateTime(day, month, year)
	leapStart := LeapStart(year)
	leapEnd := LeapEnd(year)

	if IsLeapYear(year) {
		if (d.Equal(leapStart) || d.After(leapStart)) && d.Before(leapEnd) {
			return getSaintTroparionData(gdb, d.AddDate(0, 0, 1))
		} else if d.Equal(leapEnd) {
			return getSaintTroparionData(gdb, leapStart)
		} else {
			return getSaintTroparionData(gdb, d)
		}
	} else {
		s, err := getSaintTroparionData(gdb, d)
		if err == nil && d.Equal(leapEnd) {
			s1, err := getSaintTroparionData(gdb, LeapStart(2000))
			return append(s, s1...), err
		}

		return s, err
	}
}

func GetFeastTroparion(id string) (trop []Troparion, err error) {
	dbName := "assets/troparia/troparion_feast.sqlite"
	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	s := []Troparion{}

	if result := gdb.Table("tropari").Where("comment = ?", id).Find(&s); result.Error != nil {
		return nil, result.Error
	}

	return s, nil
}
