package api

import (
	"fmt"

	"github.com/szmcdull/glinq/garray"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BookData struct {
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Code        string     `json:"code"`
	ContentType string     `json:"contentType"`
	Sections    []string   `json:"sections"`
	Items       [][]string `json:"items"`
}

type BookTitle struct {
}

func loadString(gdb *gorm.DB, key string) (value string, err error) {
	result := map[string]interface{}{}

	if err := gdb.Table("data").Find(&result, "key = ?", key).Order("key"); err.Error != nil {
		return "", err.Error
	}

	if value, ok := result["value"].(string); ok {
		return value, nil

	} else {
		return "", nil
	}

}

func GetBookData(filename string) (data BookData, err error) {
	dbName := fmt.Sprintf("assets/books/%s", filename)

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return BookData{}, err
	}

	title, err1 := loadString(gdb, "title")
	author, err2 := loadString(gdb, "author")
	code, err3 := loadString(gdb, "code")
	contentType, err4 := loadString(gdb, "contentType")

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return BookData{}, err
	}

	var sections []struct{ Title string }

	if result := gdb.Table("sections").Order("id").Find(&sections); result.Error != nil {
		return BookData{}, result.Error
	}

	sec_titles := garray.MapI(sections, func(i int) string { return sections[i].Title })

	item_titles := garray.MapI(sections, func(i int) []string {
		var items []struct{ Title string }
		if result := gdb.Table("content").Where("section = ?", i).Order("item").Find(&items); result.Error != nil {
			return []string{}
		}

		return garray.MapI(items, func(j int) string { return items[j].Title })
	})

	return BookData{Title: title, Author: author, Code: code, ContentType: contentType, Sections: sec_titles, Items: item_titles}, nil
}
