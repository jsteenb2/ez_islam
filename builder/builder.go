package builder

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/jsteenb2/quran/model"
)

func GenerateQuran(db model.DBface, baseURL, repoPath string) {
	sahihEdition, err := model.GetQuran([]byte("quran"), []byte("en.sahih"), db)
	if err != nil {
		panic(err)
	}

	CreateHTMLFiles(sahihEdition, baseURL, repoPath)
}

func CreateHTMLFiles(quranEdition model.QuranMeta, baseURL, repoPath string) {
	templates, err := template.ParseGlob("templates/*.tmpl")
	checkLog(err)
	pathPrefix := strings.Join([]string{repoPath, "public", quranEdition.Identifier}, "/")

	for idx := range quranEdition.Suwar {
		CreateSurahHTMLFile(pathPrefix, baseURL, quranEdition.Suwar[idx], templates)
	}
}
func CreateSurahHTMLFile(pathPrefix, baseURL string, surah model.SuraMeta, templates *template.Template) {
	path := fmt.Sprintf("%s/%d-%s/", pathPrefix, surah.Number, strings.ToLower(surah.EnglishName))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		checkLog(err)
	}
	outputFile, err := os.Create(path + "index.html")
	if hasError := checkLog(err); hasError == true {
		return
	}
	defer outputFile.Close()
	surahContext := SurahContext{surah, baseURL}
	templates.ExecuteTemplate(outputFile, "content.tmpl", surahContext)
	fmt.Println("successfully created ", surah.EnglishName)
}

type SurahContext struct {
	model.SuraMeta
	BaseURL string
}

func checkLog(err error) (hasError bool) {
	if err != nil {
		log.Println("E! ", err)
		hasError = true
	}
	hasError = false
	return
}
