package builder

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

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

	pathPrefix := filepath.Join(repoPath, "public", quranEdition.Identifier)
	for idx := range quranEdition.Suwar {
		surah := NewSurahContext(idx, quranEdition.Identifier, baseURL, quranEdition)
		CreateSurahHTMLFile(pathPrefix, surah, templates)
	}
}
func CreateSurahHTMLFile(pathPrefix string, surah SurahContext, templates *template.Template) {
	path := filepath.Join(pathPrefix, surah.pathID())
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	outputFile, createErr := os.Create(filepath.Join(path, "index.html"))
	if hasError := checkLog(createErr); hasError {
		return
	}

	defer outputFile.Close()
	if tempErr := templates.ExecuteTemplate(outputFile, "content.tmpl", surah); !checkLog(tempErr) {
		fmt.Println("successfully created ", surah.SurahWrap.EnglishName)
	}
}

func checkLog(err error) (hasError bool) {
	if err != nil {
		log.Println("E! ", err)
		hasError = true
	}
	hasError = false
	return
}
