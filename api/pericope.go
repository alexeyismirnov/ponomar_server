package api

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Pericope struct {
	BookName  string `json:"bookName"`
	Lang      string `json:"lang"`
	WhereExpr string `json:"whereExpr"`
}

type BibleVerse struct {
	Verse int    `json:"verse"`
	Text  string `json:"text"`
}

// curl  -X POST http://127.0.0.1:8080/pericope \
//  -d '{ "lang": "ru", "bookName": "1cor", "whereExpr": "chapter=1 AND verse>=26 AND verse<=29" }'

func GetPericope(params *Pericope) ([]BibleVerse, error) {
	dbName := fmt.Sprintf("assets/bible/%s_%s.sqlite", params.BookName, params.Lang)

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	verses := []BibleVerse{}

	if result := gdb.Table("scripture").Where(params.WhereExpr).Order("verse").Find(&verses); result.Error != nil {
		return nil, result.Error
	}

	return verses, nil
}
