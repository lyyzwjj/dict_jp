package model

type WordType int64

const (
	Sentence     = WordType(1 << 0)  //	句子
	Noun         = WordType(1 << 1)  //	名		名
	Verb1        = WordType(1 << 2)  //	動Ⅰ		动Ⅰ
	Verb2        = WordType(1 << 3)  //	動Ⅱ		动Ⅱ
	Verb3        = WordType(1 << 4)  //	動Ⅲ		动Ⅲ
	AdjectiveI   = WordType(1 << 5)  //	イ形		い形
	AdjectiveNa  = WordType(1 << 6)  //	ナ形		な形
	Auxiliary    = WordType(1 << 8)  //	助詞		助词
	Pronoun      = WordType(1 << 9)  //	代名詞	代词
	Interjection = WordType(1 << 10) //	感		感叹词
	Adverb1      = WordType(1 << 11) //	副		副
	Adverb2      = WordType(1 << 12) //	副詞		副词
	Conjunction1 = WordType(1 << 13) //	連語		连语
	Conjunction2 = WordType(1 << 14) //	連体		连体
	Conjunction3 = WordType(1 << 15) //	接		接
	Conjunction4 = WordType(1 << 16) //	接詞		接词
	Conjunction5 = WordType(1 << 17) //	接尾		接尾
	Build        = WordType(1 << 18) //			造语
	Quantifier   = WordType(1 << 19) //			助数
)

type BookVolume int8

const (
	PrimaryVolumeOne      = BookVolume(1)
	PrimaryVolumeTwo      = BookVolume(2)
	IntermediateVolumeOne = BookVolume(3)
	IntermediateVolumeTwo = BookVolume(4)
)

type WordRelation struct {
	Id    int
	Word1 Word
	Word2 Word
}

type Vocabulary struct {
	Id    int
	Kana  string  // かな 假名
	Kanji *string // 漢字 汉字
}

type WordMeaning struct {
	WordId  int
	Meaning string
	Explain string
}

type Word struct {
	Id            int
	Kana          string  // かな 假名
	Kanji         *string // 漢字 汉字
	Original      bool
	WordRelations []WordRelation
	WordMeanings  []WordMeaning
	WordType      WordType
	BookVolume    *BookVolume
}
