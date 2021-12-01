package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/lyyzwjj/dict_jp/dao"
	k "github.com/lyyzwjj/kana"
	"gorm.io/gorm"
	"os"
	"path"
	"strings"
	"testing"
)

//var (
//	word  Word
//	words []Word
//)

func TestWordInsert(t *testing.T) {
	words := []Word{
		{
			WordCore: WordCore{
				Kana:  "〜ベん",
				Kanji: "〜弁",
			},
			WordTypeValue: WordTypeConjunction4.value,
			PitchAccent:   PitchAccentLvNil,
			Meaning:       "~方言",
			Description:   "大阪弁(おおさかべん):大阪话、大阪方言",
			WordBooks: []WordBook{
				{
					Book:   BookPrimaryTwo,
					UnitNo: 26,
				},
			},
		},
		{
			WordCore: WordCore{
				Kana:  "エドヤストア",
				Kanji: "江戸屋ストア",
			},
			WordTypeValue: WordTypeConjunction4.value,
			PitchAccent:   PitchAccentLv4,
			Meaning:       "江户屋(虚构的商店)",
			WordBooks: []WordBook{
				{
					Book:   BookPrimaryTwo,
					UnitNo: 26,
				},
			},
		},
		{
			WordCore: WordCore{
				Kana:  "おくれる",
				Kanji: "遅れる",
			},
			WordTypeValue: WordTypeVerb2.value,
			PitchAccent:   PitchAccentLv0,
			Meaning:       "晚了、没赶上",
			WordBooks: []WordBook{
				{
					Book:   BookPrimaryTwo,
					UnitNo: 26,
				},
			},
		},
	}
	//for _, w := range words {
	//	w.Romaji = k.KanaToRomaji(w.Kana)
	//}
	for index, w := range words {
		words[index].Romaji = k.KanaToRomaji(w.Kana)
		// dao.Repo.Create(&words[index])
	}
	dao.Repo.Create(&words)
}

func TestWordSelect(t *testing.T) {
	var words []Word
	dao.Repo.Where("kana LIKE ?", "%か%").Find(&words)
	fmt.Println(words)
}
func TestDataInsert(t *testing.T) {
	wd := WordData{
		Kana:           "〜ベん",
		Kanji:          "〜弁",
		WordTypeName:   WordTypeConjunction4.name,
		PitchAccent:    PitchAccentLvNil,
		Meaning:        "~方言",
		Description:    "大阪弁(おおさかべん):大阪话、大阪方言",
		Masu:           "",
		TransitiveType: TransitiveTypeNil,
		Preposition:    "",
		Book:           BookPrimaryTwo,
		UnitNo:         26,
	}
	dataInsert(&wd)
	wd.UnitNo = 44
	dataInsert(&wd)
	wd.UnitNo = 26
	dataDelete(&wd)
}

func dataDelete(wd *WordData) {
	var word Word
	if err := dao.Repo.Where("Kana = ? AND Kanji >= ?", wd.Kana, wd.Kanji).Preload("WordBooks").First(&word).Error; err == nil {
		// dao.Repo.Select("WordBooks").Where("unit_no != ?", wd.UnitNo).Delete(&Word{})
		// dao.Repo.Model(&word).Association("WordBooks").Clear()
		//newWordBooks := make([]WordBook, 1)
		//for _, wb := range word.WordBooks {
		//	if wb.UnitNo == wd.UnitNo {
		//		newWordBooks = append(newWordBooks, wb)
		//	}
		//}
		//word.WordBooks = newWordBooks
		//dao.Repo.Save(&word)
		dao.Repo.Where("word_id = ? and unit_no != ?", word.ID, wd.UnitNo).Delete(&WordBook{})
	}
}

func dataInsert(wd *WordData) {
	var word Word
	if err := dao.Repo.Where("Kana = ? AND Kanji >= ?", wd.Kana, wd.Kanji).Preload("WordBooks").First(&word).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("query failed err: %#v\n", err)
			return
		}
		var checkOk bool
		if word, checkOk = wd.Data2Model(); !checkOk {
			fmt.Println("Check failed insert failed")
			return
		}
		if err := dao.Repo.Create(&word).Error; err != nil {
			fmt.Printf("word insert failed err: %#v", err)
			return
		}
	} else {
		if update, checkOk := wd.DataMergeModel(&word); checkOk {
			if !update {
				fmt.Println("word doesn't update")
				return
			}
			if err := dao.Repo.Save(&word).Error; err != nil {
				fmt.Printf("word update failed err: %#v", err)
				return
			}
		} else {
			fmt.Println("Check failed update failed")
			return
		}
	}
}

func TestWriteCsv(t *testing.T) {
	WriteCsv()
	//wd := WordData{
	//	Kana:           "〜ベん",
	//	Kanji:          "〜弁",
	//	WordTypeName:   WordTypeConjunction4.name,
	//	PitchAccent:    PitchAccentLvNil,
	//	Meaning:        "~方言",
	//	Description:    "大阪弁(おおさかべん):大阪话、大阪方言",
	//	Masu:           "",
	//	TransitiveType: TransitiveTypeNil,
	//	Preposition:    "",
	//	Book:           BookPrimaryTwo,
	//	UnitNo:         26,
	//}
	//str := wd.WordDataCsvStringer()
	//fmt.Println(str)
	//parser := WordDataCsvParser(str)
	//fmt.Printf("%#v\n", parser)
}
func TestReadAllCsv(t *testing.T) {
	ReadAllCsv()
}
func ReadAllCsv() {
	csvFilePath := "resources/raw/大家的日语第二版初级2_26.csv"
	fileNameWithoutExt := strings.TrimSuffix(path.Base(csvFilePath), path.Ext(csvFilePath))
	parts := strings.Split(fileNameWithoutExt, "_")
	if len(parts) != 2 {
		fmt.Println("fileName err", fileNameWithoutExt)
		return
	}
	book := parts[0]
	unitNo := parts[1]
	f, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Open csv file failed Error: ", err)
		return
	}
	reader := csv.NewReader(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Close csv file failed Error: ", err)
		}
	}(f)
	// 可以一次性读完
	arrs, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Read csv file failed Error: ", err)
		return
	}
	for _, arr := range arrs {
		arr = arr[0 : len(arr)-1]
		arr = append(arr, book, unitNo)
		wd := WordDataCsvParser(arr)
		dataInsert(&wd)
	}
}

func WriteCsv() {
	csvFilePath := "resources/raw/test.csv"
	f, err := os.OpenFile(csvFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	defer writer.Flush()

	wd := WordData{
		Kana:           "〜ベん",
		Kanji:          "〜弁",
		WordTypeName:   WordTypeConjunction4.name,
		PitchAccent:    PitchAccentLvNil,
		Meaning:        "~方言",
		Description:    "大阪弁(おおさかべん):大阪话、大阪方言",
		Masu:           "",
		TransitiveType: TransitiveTypeNil,
		Preposition:    "",
		Book:           BookPrimaryTwo,
		UnitNo:         26,
	}
	arr := wd.WordDataCsvStringer()
	fmt.Println(arr)
	// writer.WriteAll([][]string{arr})
	if err := writer.Write(arr); err != nil {
		fmt.Println("write to csv file failed", arr)
		return
	}
	//t := reflect.TypeOf(&WordData{})
	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//}
	//if t.Kind() != reflect.Struct {
	//	log.Println("Check type error not Struct")
	//	return
	//}
	//fieldNum := t.NumField()
	//header := make([]string, 0, fieldNum)
	//for i := 0; i < fieldNum; i++ {
	//	header = append(header, t.Field(i).Name)
	//}
	//writer.Write(header)
	//var data = []string{"3", "John", "23"}
	//writer.Write(data)
	//// 也可以一次性写入多条
	//var d = [][]string{{"1", "Edgar", "20"}, {"2", "Tom", "18"}}
	//writer.WriteAll(d)
	//// 将缓存中的内容写入到文件里
	//writer.Flush()
	//if err = writer.Error(); err != nil {
	//	fmt.Println(err)
	//}
}
