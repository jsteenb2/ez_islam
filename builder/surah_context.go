package builder

import (
	"fmt"
	"strings"

	"github.com/jsteenb2/quran/model"
)

type SurahContext struct {
	SurahWrap
	PrevSura SurahWrap
	NextSura SurahWrap
	BaseURL  string
	Edition  string
}

func NewSurahContext(currentIndex int, edition, baseURL string, quran model.QuranMeta) (surah SurahContext) {
	surah.BaseURL = baseURL
	surah.SurahWrap = SurahWrap(quran.Suwar[currentIndex])
	surah.Edition = edition
	if currentIndex != 0 {
		surah.PrevSura = SurahWrap(quran.Suwar[currentIndex-1])
	}

	if currentIndex != 113 {
		surah.NextSura = SurahWrap(quran.Suwar[currentIndex+1])
	}
	return
}

func (s SurahContext) pathID() string {
	return fmt.Sprintf("%d-%s", s.Number, strings.ToLower(s.EnglishName))
}
