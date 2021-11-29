package model

import (
	"fmt"
	"github.com/lyyzwjj/dict_jp/dao"
	k "github.com/lyyzwjj/kana"
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
			Description:   "大阪弁(おおさかべん):大版话、大阪方言",
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
		Description:    "大阪弁(おおさかべん):大版话、大阪方言",
		Masu:           "",
		TransitiveType: TransitiveTypeNil,
		Preposition:    "",
		Book:           BookPrimaryTwo,
		UnitNo:         26,
	}
	var word Word
	if err := dao.Repo.Where("Kana = ? AND Kanji >= ?", wd.Kana, wd.Kanji).Preload("WordBooks").First(&word).Error; err != nil {
		fmt.Println(word)
	} else {
		fmt.Println(word)
	}
	if err := dao.Repo.Where("Kana = ? AND Kanji >= ?", wd.Kana, wd.Kanji).First(&word).Error; err != nil {
		var ok bool
		if word, ok = wd.Data2Model(); ok {
			dao.Repo.Create(&word)
		}
		if err := dao.Repo.Where("Kana = ? AND Kanji >= ?", wd.Kana, wd.Kanji).First(&word).Error; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(word)
		}
	} else {
		err := dao.Repo.Model(&word).Association("WordBooks").Find(&word.WordBooks)
		if err != nil {
			fmt.Print(err)
			return
		}
		if ok := wd.DataMergeModel(&word); ok {
			dao.Repo.Save(&word)
		}
		if err := dao.Repo.Where("Kana = ? AND Kanji >= ?", wd.Kana, wd.Kanji).First(&word).Error; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(word)
		}
	}

}
