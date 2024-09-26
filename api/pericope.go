package api

import "fmt"

type Pericope struct {
	BookName  string `json:"bookName"`
	Lang      string `json:"lang"`
	WhereExpr string `json:"whereExpr"`
}

type BibleVerse struct {
	Verse int    `json:"verse"`
	Text  string `json:"text"`
}

func GetPericope(params *Pericope) ([]BibleVerse, error) {
	str := fmt.Sprintf("%s - %s", params.BookName, params.WhereExpr)
	return []BibleVerse{{1, str}}, nil
}
