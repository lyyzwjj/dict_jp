package model

import (
	"fmt"
	"github.com/lyyzwjj/dict_jp/dao"
	"github.com/lyyzwjj/dict_jp/wjjutils"
	"testing"
)

var (
	word  Word
	words []Word
)

func TestWordInsert(t *testing.T) {
	wjjutils.InitSnowflake()
	dao.InitMySQLDB()
	//word := &Word{
	//	VocabularyCore: VocabularyCore{
	//		Kana: "ガム",
	//	},
	//	WordMeanings: []WordMeaning{
	//		{
	//			Major:       true,
	//			WordType:    Noun,
	//			Meaning:     "口香糖",
	//			Description: "gum",
	//			BookVolume:  PrimaryVolumeTwo,
	//			UnitNo:      28,
	//		},
	//	},
	//}
	// vocabulary := NewVocabulary(word.VocabularyCore, true)
	// dao.Repo.Create(vocabulary)
	// dao.Repo.Create(word)
	// dao.Repo.Save(word)

	word := &Word{
		VocabularyCore: VocabularyCore{
			Kana:  "かむ",
			Kanji: "噛む",
		},
		WordMeanings: []WordMeaning{
			{
				Major:       true,
				WordType:    Verb1,
				Meaning:     "嚼、咬",
				Description: "ガムを噛みます",
				BookVolume:  PrimaryVolumeTwo,
				UnitNo:      28,
			},
		},
	}
	dao.Repo.Save(word)
}

func TestWordSelect(t *testing.T) {
	dao.InitMySQLDB()
	dao.Repo.Where("kana LIKE ?", "%か%").Find(&words)
	fmt.Println(words)
}
