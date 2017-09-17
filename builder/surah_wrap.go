package builder

import (
	"fmt"
	"strings"

	"github.com/jsteenb2/quran/model"
)

func (s SurahWrap) SuraPath() (string, error) {
	urlPath := fmt.Sprintf(fmt.Sprintf("%d-%s", s.Number, strings.ToLower(s.EnglishName)))
	return urlPath, nil
}

type SurahWrap model.SuraMeta
