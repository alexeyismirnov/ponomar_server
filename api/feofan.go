package api

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FeofanParams struct {
	Id    string `json:"id"`
	Fuzzy bool   `json:"fuzzy"`
}

type Feofan struct {
	Descr string
}

func GetFeofan(params *FeofanParams) (_ string, err error) {
	dbName := "assets/books/feofan.sqlite"

	gdb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return "", err
	}

	feofan := Feofan{}
	whereExpr := ""

	if params.Fuzzy {
		whereExpr = fmt.Sprintf(`id LIKE "%s%s" AND fuzzy=1`, "%", params.Id)

	} else {
		whereExpr = fmt.Sprintf(`id="%s"`, params.Id)
	}

	if result := gdb.Table("thoughts").Where(whereExpr).First(&feofan); result.Error != nil {
		fmt.Println(result.Error)
		return "", result.Error
	}

	return feofan.Descr, nil
}
