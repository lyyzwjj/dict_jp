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

func (wd *WordData) WordDataCsvStringer() (arr []string) {
	arr = append(arr, wd.Masu, wd.Kanji, wd.Kana, strconv.Itoa(wd.PitchAccent), wd.WordTypeName, wd.Meaning, wd.Description, wd.TransitiveType, wd.Preposition, wd.Book, strconv.Itoa(int(wd.UnitNo)))
	return
}

func WordDataCsvParser(arr []string) WordData {
	var pitchAccent int
	for _, p := range strings.Split(arr[3], "|") {
		if subPa, err := strconv.Atoi(p); err == nil {
			if subPa == PitchAccentLvNil {
				pitchAccent = subPa
			} else {
				pitchAccent = pitchAccent | (1 << subPa)
			}
		} else {
			fmt.Printf("pitchAccent parser error, arr[3]: %v\n", arr[3])
			panic("pitchAccent parser error, arr[3]: err: " + err.Error())
		}
	}
	unitNo, _ := strconv.ParseUint(arr[10], 10, 8)
	return WordData{
		Masu:           arr[0],
		Kana:           arr[1],
		Kanji:          arr[2],
		PitchAccent:    pitchAccent,
		WordTypeName:   arr[4],
		Meaning:        arr[5],
		Description:    arr[6],
		TransitiveType: arr[7],
		Preposition:    arr[8],
		Book:           arr[9],
		UnitNo:         uint8(unitNo),
	}
}

func (wd *WordData) Check() (WordTypeValue int, ok bool) {
	if ok = checkTransitiveType(wd.TransitiveType); !ok {
		return
	}
	if _, isVerb := verbNameSet[wd.WordTypeName]; isVerb && strings.TrimSpace(wd.TransitiveType) == TransitiveTypeNil {
		return
	}
	if ok = checkBook(wd.Book); !ok {
		return
	}
	if ok = checkPitchAccent(wd.PitchAccent); !ok {
		return
	}
	if WordTypeValue, ok = wordTypeName2Value(wd.WordTypeName); !ok {
		return
	}
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
