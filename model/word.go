package model

type WordType int8

//动词1
//动词2
//动词3
//名词
//i形
//na形
//助词
//副词
//代词
//连体
//接尾
//连语
//接辞
//造语
//感
//接

const (
	Verb1       = WordType(1)
	Verb2       = WordType(2)
	Verb3       = WordType(3)
	Noun        = WordType(4)
	AdjectiveI  = WordType(5)
	AdjectiveNa = WordType(6)
	Adverb      = WordType(7)
	Conjunction = WordType(8)
	Suffix      = WordType(9)
	Prefix      = WordType(10)
	Sentence    = WordType(11)
)

type Word struct {
	Id       int
	WordType WordType
	// かな 假名
	kana string
	// 漢字
	kanji *string
}
