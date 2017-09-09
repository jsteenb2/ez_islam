package main

import (
	"os"

	"github.com/jsteenb2/ez_islam/builder"
	"github.com/jsteenb2/quran"
)

var BaseURL = os.Getenv("base")

func main() {
	db := quran.GetQuranDB()
	builder.GenerateQuran(db, BaseURL)
}
