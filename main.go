package main

import (
	"log"
	"os"

	"github.com/jsteenb2/ez_islam/builder"
	"github.com/jsteenb2/quran"
)

var (
	BASEURL = os.Getenv("base")
	DBPATH  = os.Getenv("DBPATH")
)

func main() {
	db := quran.GetQuranDB(getDBPath(DBPATH))
	builder.GenerateQuran(db, BASEURL)
}

func getDBPath(userInput string) string {
	if userInput != "" {
		return userInput
	}
	dbPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dbPath + "/quran.db"
}
