package builder

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/jsteenb2/quran/model"
)

func GenerateQuran(db model.DBface, baseURL string) {
	sahihEdition, err := model.GetQuran([]byte("quran"), []byte("en.sahih"), db)
	if err != nil {
		panic(err)
	}

	CreateHTMLFiles(sahihEdition, baseURL)
}

func CreateHTMLFiles(quranEdition model.QuranMeta, baseURL string) {
	templates, err := template.ParseGlob("templates/*.tmpl")
	checkLog(err)
	pathPrefix := fmt.Sprintf("%s/src/github.com/jsteenb2/ez_islam/public/%s", os.Getenv("GOPATH"), quranEdition.Identifier)

	for idx := range quranEdition.Suwar {
		CreateSurahHTMLFile(pathPrefix, baseURL, quranEdition.Suwar[idx], templates)
	}
}
func CreateSurahHTMLFile(pathPrefix, baseURL string, surah model.SuraMeta, templates *template.Template) {
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
