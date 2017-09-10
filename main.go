package main

import (
	"log"
	"os"

	"github.com/jsteenb2/ez_islam/builder"
	"github.com/jsteenb2/quran"
)

var (
	BASEURL  = os.Getenv("BASEURL")
	REPOPATH = os.Getenv("REPO")
	DB       = os.Getenv("DB")
)

func main() {
	db := quran.GetQuranDB(getDBPath(DB))
	builder.GenerateQuran(db, BASEURL, REPOPATH)
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
