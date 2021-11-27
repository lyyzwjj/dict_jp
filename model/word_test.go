package model

import (
	"fmt"
	"github.com/lyyzwjj/dict_jp/dao"
	"github.com/lyyzwjj/dict_jp/wjjutils"
	k "github.com/lyyzwjj/kana"
	"testing"
)

var (
	word  Word
	words []Word
)

func TestWordInsert(t *testing.T) {
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
	//word := &Word{
	//	VocabularyCore: VocabularyCore{
	//		ID:    wjjutils.GenID(),
	//		Kana:  "かむ",
	//		Kanji: "噛む",
	//	},
	//	WordMeanings: []WordMeaning{
	//		{
	//			ID:          wjjutils.GenID(),
	//			Major:       true,
	//			WordType:    Verb1,
	//			Meaning:     "嚼、咬",
	//			Description: "ガムを噛みます",
	//			BookVolume:  PrimaryVolumeTwo,
	//			UnitNo:      28,
	//		},
	//	},
	//}
	//word = Word{
	//	VocabularyCore: VocabularyCore{
	//		ID:     wjjutils.GenID(),
	//		Kana:   "〜ベん",
	//		Kanji:  "〜弁",
	//		Romaji: k.KanaToRomaji("〜ベん"),
	//	},
	//	WordMeanings: []WordMeaning{
	//		{
	//			ID:         wjjutils.GenID(),
	//			Major:      true,
	//			WordType:   Conjunction4,
	//			Meaning:    "~方言",
	//			BookVolume: PrimaryVolumeTwo,
	//			UnitNo:     26,
	//		},
	//	},
	//}
	words = []Word{
		{
			VocabularyCore: VocabularyCore{
				ID:    wjjutils.GenID(),
				Kana:  "〜ベん",
				Kanji: "〜弁",
			},
			WordType:   Conjunction4,
			Meaning:    "~方言",
			BookVolume: PrimaryVolumeTwo,
			UnitNo:     26,
		},
		{
			VocabularyCore: VocabularyCore{
				ID:    wjjutils.GenID(),
				Kana:  "おおさかべん",
				Kanji: "大阪弁",
			},
			WordType:   Conjunction4,
			Meaning:    "大版话、大阪方言",
			BookVolume: PrimaryVolumeTwo,
			UnitNo:     26,
		},
		{
			VocabularyCore: VocabularyCore{
				ID:    wjjutils.GenID(),
				Kana:  "エドヤストア",
				Kanji: "江戸屋ストア",
			},
			PitchAccent: PitchAccentLv4,
			WordType:    Conjunction4,
			Meaning:     "江户屋(虚构的商店)",
			BookVolume:  PrimaryVolumeTwo,
			UnitNo:      26,
		},
		{
			VocabularyCore: VocabularyCore{
				ID:    wjjutils.GenID(),
				Kana:  "おくれる",
				Kanji: "遅れる",
			},
			PitchAccent: PitchAccentLv0,
			WordType:    Verb2,
			Meaning:     "晚了、没赶上",
			BookVolume:  PrimaryVolumeTwo,
			UnitNo:      26,
		},
	}
	//for _, w := range words {
	//	w.Romaji = k.KanaToRomaji(w.Kana)
	//}
	for index, w := range words {
		words[index].Romaji = k.KanaToRomaji(w.Kana)
	}
	dao.Repo.Create(&words)
}

func TestWordSelect(t *testing.T) {
	dao.InitMySQLDB()
	dao.Repo.Where("kana LIKE ?", "%か%").Find(&words)
	fmt.Println(words)
}
