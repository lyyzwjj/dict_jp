package model

type WordType int64

//	句子		句子
//	名詞		名词
//	動詞Ⅰ	动词Ⅰ
//	動詞Ⅱ	动词Ⅱ
//	動詞Ⅲ	动词Ⅲ
//	イ形		い形
//	ナ形		な形
//	助詞		助词
//	代名詞	代词
//	感
//	副
//	副詞		副词

//	连语		前
//	连体
//	接		后
//	接辞
//	接尾
//	造语
//	助数

const (
	Sentence     = WordType(1 << 0)  //	句子
	Noun         = WordType(1 << 1)  //	名詞		名词
	Verb1        = WordType(1 << 2)  //	動詞Ⅰ	动词Ⅰ
	Verb2        = WordType(1 << 3)  //	動詞Ⅱ	动词Ⅱ
	Verb3        = WordType(1 << 4)  //	動詞Ⅲ	动词Ⅲ
	AdjectiveI   = WordType(1 << 5)  //	イ形		い形
	AdjectiveNa  = WordType(1 << 6)  //	ナ形		な形
	Auxiliary    = WordType(1 << 8)  //	助詞		助词
	Pronoun      = WordType(1 << 9)  //	代名詞	代词
	Interjection = WordType(1 << 10) //	感		感叹词
	Adverb1      = WordType(1 << 11) //	副		副
	Adverb2      = WordType(1 << 12) //	副詞		副词
	Conjunction1 = WordType(1 << 13) //			连语
	Conjunction2 = WordType(1 << 14) //			连体
	Prefix       = WordType(1 << 15) //			接
	Prefix       = WordType(1 << 16) //			接词
	Prefix       = WordType(1 << 17) //			接尾
	Sentence     = WordType(1 << 18) //			造语
	Sentence     = WordType(1 << 19) //			助数
)

type Word struct {
	Id       int
	WordType WordType
	// かな 假名
	Kana string
	// 漢字 汉字
	Kanji *string
	//

}
