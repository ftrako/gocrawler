package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/db"
)

type IFileParser interface {
	StartUrl() string
	FileFilter(url string) bool
	FileParse(doc *goquery.Document, db *db.FileDB)
}
