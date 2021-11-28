package model

import (
	"github.com/lyyzwjj/kana"
	"strings"
)

type WordData struct {
	Kana           string
	Kanji          string
	WordTypeName   string
	PitchAccent    int
	Meaning        string
	Description    string
	Masu           string
	TransitiveType string
	Preposition    string
	Book           string
	UnitNo         uint8
}

func (wd *WordData) Check() (WordTypeValue int, ok bool) {
	if ok = checkTransitiveType(wd.TransitiveType); !ok {
		return
	}
	if ok = checkBook(wd.Book); !ok {
		return
	}
	if ok = checkPitchAccent(wd.PitchAccent); !ok {
		return
	}
	WordTypeValue, ok = wordTypeName2Value(wd.WordTypeName)
	return
}

func (wd *WordData) Data2Model() (word Word, ok bool) {
	var WordTypeValue int
	if WordTypeValue, ok = wd.Check(); ok {
		word = Word{
			WordCore: WordCore{
				Kana:   wd.Kana,
				Kanji:  wd.Kanji,
				Romaji: kana.KanaToRomaji(wd.Kana),
			},
			WordTypeValue:  WordTypeValue,
			PitchAccent:    wd.PitchAccent,
			Meaning:        wd.Meaning,
			Description:    wd.Description,
			Masu:           wd.Masu,
			TransitiveType: wd.TransitiveType,
			Preposition:    wd.Preposition,
			WordBooks: []WordBook{
				{
					Book:   wd.Book,
					UnitNo: wd.UnitNo,
				},
			},
		}
	}
	return
}

func (wd *WordData) DataMergeModel(w *Word) (result bool) {
	var wordTypeValueFrom int
	if wordTypeValueFrom, result = wd.Check(); result {
		if wordTypeValue, ok := mergeWordType(wordTypeValueFrom, w.WordTypeValue); ok {
			result = true
			w.WordTypeValue = wordTypeValue
		}
		if pitchAccent, ok := mergePitchAccent(wd.PitchAccent, w.PitchAccent); ok {
			result = true
			w.PitchAccent = pitchAccent
		}
		trimMeaning := strings.TrimSpace(wd.Meaning)
		if trimMeaning != "" && !strings.Contains(w.Meaning, trimMeaning) {
			result = true
			w.Meaning = w.Meaning + "\n" + wd.Meaning
		}
		trimDescription := strings.TrimSpace(wd.Description)
		if trimDescription != "" && !strings.Contains(w.Description, trimDescription) {
			result = true
			w.Description = w.Description + "\n" + wd.Description
		}
		if isVerbByName(wd.WordTypeName) && !isVerbByValue(w.WordTypeValue) {
			masu := strings.TrimSpace(wd.Masu)
			if masu != "" && w.Masu != masu {
				result = true
				w.Masu = masu
			}
			preposition := strings.TrimSpace(wd.Preposition)
			if preposition != "" && w.Preposition != preposition {
				result = true
				w.Preposition = preposition
			}
		}
		existWordBook := false
		for _, wordBook := range w.WordBooks {
			if wordBook.Book == wd.Book && wordBook.UnitNo == wd.UnitNo {
				existWordBook = true
				break
			}
		}
		if !existWordBook {
			result = true
			w.WordBooks = append(w.WordBooks, WordBook{
				Book:   wd.Book,
				UnitNo: wd.UnitNo,
			})
		}
	}
	return
}
