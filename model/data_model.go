package model

import (
	"fmt"
	"github.com/lyyzwjj/kana"
	"strconv"
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

func (wd *WordData) WordDataCsvStringer() string {
	return fmt.Sprintf("%s,%s,%s,%d,%s,%s,%s,%s,%s,%s,%d", wd.Kana, wd.Kanji, wd.WordTypeName, wd.PitchAccent, wd.Meaning, wd.Description, wd.Masu, wd.TransitiveType, wd.Preposition, wd.Book, wd.UnitNo)
}

func WordDataCsvParser(line string) WordData {
	fields := strings.Split(line, ",")
	pitchAccent, _ := strconv.Atoi(fields[3])
	unitNo, _ := strconv.ParseUint(fields[10], 10, 8)
	return WordData{
		Kana:           fields[0],
		Kanji:          fields[1],
		WordTypeName:   fields[2],
		PitchAccent:    pitchAccent,
		Meaning:        fields[4],
		Description:    fields[5],
		Masu:           fields[6],
		TransitiveType: fields[7],
		Preposition:    fields[8],
		Book:           fields[9],
		UnitNo:         uint8(unitNo),
	}
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

func (wd *WordData) Data2Model() (word Word, checkOk bool) {
	var WordTypeValue int
	if WordTypeValue, checkOk = wd.Check(); checkOk {
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

func (wd *WordData) DataMergeModel(w *Word) (update, checkOk bool) {
	var wordTypeValueFrom int
	if wordTypeValueFrom, checkOk = wd.Check(); checkOk {
		if wordTypeValue, ok := mergeWordType(wordTypeValueFrom, w.WordTypeValue); ok {
			update = true
			w.WordTypeValue = wordTypeValue
		}
		if pitchAccent, ok := mergePitchAccent(wd.PitchAccent, w.PitchAccent); ok {
			update = true
			w.PitchAccent = pitchAccent
		}
		trimMeaning := strings.TrimSpace(wd.Meaning)
		if trimMeaning != "" && !strings.Contains(w.Meaning, trimMeaning) {
			update = true
			w.Meaning = w.Meaning + "\n" + wd.Meaning
		}
		trimDescription := strings.TrimSpace(wd.Description)
		if trimDescription != "" && !strings.Contains(w.Description, trimDescription) {
			update = true
			w.Description = w.Description + "\n" + wd.Description
		}
		if isVerbByName(wd.WordTypeName) && !isVerbByValue(w.WordTypeValue) {
			masu := strings.TrimSpace(wd.Masu)
			if masu != "" && w.Masu != masu {
				update = true
				w.Masu = masu
			}
			preposition := strings.TrimSpace(wd.Preposition)
			if preposition != "" && w.Preposition != preposition {
				update = true
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
			update = true
			w.WordBooks = append(w.WordBooks, WordBook{
				Book:   wd.Book,
				UnitNo: wd.UnitNo,
			})
		}
	}
	return
}
