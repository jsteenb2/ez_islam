package main

import (
	"os"

	"github.com/jsteenb2/ez_islam/builder"
	"github.com/jsteenb2/quran"
	"log"
)

var BaseURL = os.Getenv("base")

func main() {
	db := quran.GetQuranDB(getDBPath())
	builder.GenerateQuran(db, BaseURL)
}

func getDBPath() string {
	dbPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dbPath + "/quran.db"
}
