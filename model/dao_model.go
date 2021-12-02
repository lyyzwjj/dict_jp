package model

type wordType struct {
	name  string
	value int
}

const (
	BookNil                    = ""
	BookPrimaryOne             = "大家的日语第二版初级1"
	BookPrimaryTwo             = "大家的日语第二版初级2"
	BookIntermediateOne        = "大家的日语第二版中级1"
	BookIntermediateTwo        = "大家的日语第二版中级2"
	PitchAccentLvNil           = -1
	PitchAccentLv0             = 1 << 0
	PitchAccentLv1             = 1 << 1
	PitchAccentLv2             = 1 << 2
	PitchAccentLv3             = 1 << 3
	PitchAccentLv4             = 1 << 4
	PitchAccentLv5             = 1 << 5
	PitchAccentLv6             = 1 << 6
	PitchAccentLv7             = 1 << 7
	PitchAccentLv8             = 1 << 8
	PitchAccentLv9             = 1 << 9
	PitchAccentLvAll           = PitchAccentLv0 | PitchAccentLv1 | PitchAccentLv2 | PitchAccentLv3 | PitchAccentLv4 | PitchAccentLv5 | PitchAccentLv6 | PitchAccentLv7 | PitchAccentLv8 | PitchAccentLv9
	TransitiveTypeNil          = ""
	TransitiveTypeBoth         = "自他"
	TransitiveTypeTransitive   = "他"
	TransitiveTypeIntransitive = "自"
)

var (
	WordTypeNil          = wordType{"", 1 << 0}
	WordTypeVerb1        = wordType{"動Ⅰ", 1 << 1}  //	動Ⅰ		动Ⅰ
	WordTypeVerb2        = wordType{"動Ⅱ", 1 << 2}  //	動Ⅱ		动Ⅱ
	WordTypeVerb3        = wordType{"動Ⅲ", 1 << 3}  //	動Ⅲ		动Ⅲ
	WordTypeNoun         = wordType{"名", 1 << 4}   //	名		名
	WordTypeAdjectiveI   = wordType{"イ形", 1 << 5}  //	イ形		い形
	WordTypeAdjectiveNa  = wordType{"ナ形", 1 << 6}  //	ナ形		な形
	WordTypeAuxiliary    = wordType{"助", 1 << 7}   //	助		助
	WordTypePronoun      = wordType{"代", 1 << 8}   //	代名詞	代词
	WordTypeInterjection = wordType{"感", 1 << 9}   //	感		感叹词
	WordTypeAdverb1      = wordType{"副", 1 << 10}  //	副		副
	WordTypeAdverb2      = wordType{"副詞", 1 << 11} //	副詞		副词
	WordTypeConjunction1 = wordType{"連語", 1 << 12} //	連語		连语
	WordTypeConjunction2 = wordType{"連体", 1 << 13} //	連体		连体
	WordTypeConjunction3 = wordType{"接", 1 << 14}  //	接		接
	WordTypeConjunction4 = wordType{"接辞", 1 << 15} //	接辞		接词
	WordTypeConjunction5 = wordType{"接尾", 1 << 16} //	接尾		接尾
	WordTypeBuildLang    = wordType{"造語", 1 << 17} //			造语
	WordTypeQuantifier   = wordType{"助数", 1 << 18} //			助数
	verbNameSet          = map[string]wordType{
		WordTypeVerb1.name: WordTypeVerb1,
		WordTypeVerb2.name: WordTypeVerb2,
		WordTypeVerb3.name: WordTypeVerb3,
	}
	verbValueSet = map[int]wordType{
		WordTypeVerb1.value: WordTypeVerb1,
		WordTypeVerb2.value: WordTypeVerb2,
		WordTypeVerb3.value: WordTypeVerb3,
	}
	wordTypeSet = map[string]wordType{
		WordTypeNil.name:          WordTypeNil,
		WordTypeNoun.name:         WordTypeNoun,
		WordTypeVerb1.name:        WordTypeVerb1,
		WordTypeVerb2.name:        WordTypeVerb2,
		WordTypeVerb3.name:        WordTypeVerb3,
		WordTypeAdjectiveI.name:   WordTypeAdjectiveI,
		WordTypeAdjectiveNa.name:  WordTypeAdjectiveNa,
		WordTypeAuxiliary.name:    WordTypeAuxiliary,
		WordTypePronoun.name:      WordTypePronoun,
		WordTypeInterjection.name: WordTypeInterjection,
		WordTypeAdverb1.name:      WordTypeAdverb1,
		WordTypeAdverb2.name:      WordTypeAdverb2,
		WordTypeConjunction1.name: WordTypeConjunction1,
		WordTypeConjunction2.name: WordTypeConjunction2,
		WordTypeConjunction3.name: WordTypeConjunction3,
		WordTypeConjunction4.name: WordTypeConjunction4,
		WordTypeConjunction5.name: WordTypeConjunction5,
		WordTypeBuildLang.name:    WordTypeBuildLang,
		WordTypeQuantifier.name:   WordTypeQuantifier,
	}

	transitiveTypeSet = map[string]bool{
		TransitiveTypeNil:          true,
		TransitiveTypeBoth:         true,
		TransitiveTypeTransitive:   true,
		TransitiveTypeIntransitive: true,
	}
	bookSet = map[string]bool{
		BookNil:             true,
		BookPrimaryOne:      true,
		BookPrimaryTwo:      true,
		BookIntermediateOne: true,
		BookIntermediateTwo: true,
	}
)

func wordTypeName2Value(wordTypeFromName string) (wordTypeValue int, ok bool) {
	var wordTypeFrom wordType
	if wordTypeFrom, ok = wordTypeSet[wordTypeFromName]; ok {
		wordTypeValue = wordTypeFrom.value
	}
	return
}
func checkPitchAccent(pitchAccentFrom int) (ok bool) {
	ok = pitchAccentFrom == PitchAccentLvNil || pitchAccentFrom&PitchAccentLvAll == pitchAccentFrom
	return
}

func checkTransitiveType(transitive string) (ok bool) {
	_, ok = transitiveTypeSet[transitive]
	return
}
func checkBook(book string) (ok bool) {
	_, ok = bookSet[book]
	return
}
func isVerbByName(verbTypeName string) (ok bool) {
	_, ok = verbNameSet[verbTypeName]
	return
}
func isVerbByValue(verbTypeValue int) (ok bool) {
	for key := range verbValueSet {
		if ok = verbTypeValue&key == key; ok {
			break
		}
	}
	return
}

//func mergeWordType(wordTypeFromName string, wordTypeToValue int) (wordTypeValue int, ok bool) {
//	var wordTypeFrom wordType
//	if wordTypeFrom, ok = wordTypeSet[wordTypeFromName]; ok {
//		wordTypeValue = wordTypeFrom.value & wordTypeToValue
//	}
//	return
//}

func mergeWordType(wordTypeFromValue, wordTypeToValue int) (wordTypeValue int, ok bool) {
	if ok = wordTypeFromValue != wordTypeToValue; ok {
		wordTypeValue = wordTypeFromValue | wordTypeToValue
	}
	return
}

func mergePitchAccent(pitchAccentFrom, pitchAccentTo int) (pitchAccent int, ok bool) {
	if ok = pitchAccentFrom != pitchAccentTo; ok {
		if pitchAccentTo == PitchAccentLvNil {
			pitchAccent = pitchAccentFrom
		} else if pitchAccentFrom == PitchAccentLvNil {
			pitchAccent = pitchAccentTo
		} else { //pitchAccentFrom != PitchAccentLvNil pitchAccentTo != PitchAccentLvNil
			pitchAccent = pitchAccentFrom | pitchAccentTo
		}
	}
	return
}

// type RelationType uint8
//type WordRelation struct {
//	ID           int
//	WordA        Word `gorm:"primaryKey"`
//	WordB        Word `gorm:"primaryKey"`
//	RelationType RelationType
//}

type WordCore struct {
	ID     uint   `gorm:"type:int(11) auto_increment;primaryKey;"`
	Kana   string `gorm:"type:varchar(255);not null;comment:かな;index:idx_Kana_kanji,unique;"` // かな 假名
	Kanji  string `gorm:"type:varchar(255);not null;comment:漢字;index:idx_Kana_kanji,unique;"` // 漢字 汉字
	Romaji string `gorm:"type:varchar(255);not null;comment:ロマ字;"`
}

type Vocabulary struct {
	WordCore
	Original bool `gorm:"default:true;not null;comment:是否是单词原形;"`
}

//func NewVocabulary(wordCore WordCore, original bool) *Vocabulary {
//	return &Vocabulary{
//		WordCore: wordCore,
//		Original: original,
//	}
//}

type WordBook struct {
	ID     uint   `gorm:"type:int(11) auto_increment;primaryKey;autoIncrement:true;"`
	WordID uint   `gorm:"type:int(11);not null;comment:关联Word表主键;"`
	Book   string `gorm:"type:varchar(20);not null;comment:教材;"`
	UnitNo uint8  `gorm:"not null;default:0;comment:课程序号;"`
}

type Word struct {
	WordCore
	WordTypeValue  int    `gorm:"type:int;not null;comment:单词类型;"`
	PitchAccent    int    `gorm:"type:int;not null;default:-1;comment:音调;"`
	Meaning        string `gorm:"type:varchar(2000);comment:释义;"`
	Description    string `gorm:"type:varchar(2000);comment:解释;"`
	TransitiveType string `gorm:"type:varchar(2);not null;comment:自他形,动词专有;"` // 自动词 他动词 自他动词
	Masu           string `gorm:"type:varchar(255);not null;comment:ます形,动词专有;"`
	Preposition    string `gorm:"type:varchar(2);not null;comment:固定搭配,特殊前置动词,动词专有;"` //自动词非が 他动词非を
	// 普通额外参数
	QueryCount uint       `gorm:"not null;default:0;comment:查询次数;"`
	WordBooks  []WordBook // `gorm:"foreignKey:WordID;references:ID"`
	// WordRelations []*Word   `gorm:"many2many:word_relations"`
}
