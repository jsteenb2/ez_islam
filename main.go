package main

import (
	"html/template"
	"log"
	"os"

	"github.com/jsteenb2/quran"
	"github.com/jsteenb2/quran/model"
	"fmt"
	"strings"
)

type Context struct {
	model.QuranMeta
	BaseURL string
}

var baseURL = os.Getenv("base")

func main() {
	db := quran.GetQuranDB()
	sahihEdition, err := model.GetQuran([]byte("quran"), []byte("en.sahih"), db)
	if err != nil {
		panic(err)
	}
	quranContext := Context{sahihEdition, ""}
	createHTMLFiles(quranContext)
}

func createHTMLFiles(quranContext Context) {
	templates, err := template.ParseGlob("templates/*.tmpl")
	checkLog(err)
	pathPrefix := fmt.Sprintf("%s/src/github.com/jsteenb2/ez_islam/public/%s", os.Getenv("GOPATH"), quranContext.Identifier)

	for idx := range quranContext.Suwar {
		createSurahHTMLFile(pathPrefix, quranContext.Suwar[idx], templates)
	}
}
func createSurahHTMLFile(pathPrefix string, surah model.SuraMeta, templates *template.Template) {
	path := fmt.Sprintf("%s/%d-%s/", pathPrefix, surah.Number, strings.ToLower(surah.EnglishName))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	outputFile, err := os.Create(path + "index.html")
	checkLog(err)
	defer outputFile.Close()
	surahContext := SurahContext{surah, baseURL}
	templates.ExecuteTemplate(outputFile, "content.tmpl", surahContext)
}

type SurahContext struct {
	model.SuraMeta
	BaseURL string
}

func checkLog(err error) {
	if err != nil {
		log.Println("E! ", err)
	}
}
