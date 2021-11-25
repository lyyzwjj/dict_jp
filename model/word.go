package model

type WordType int8

const (
	Sentence     = WordType(0)  //	句子
	Noun         = WordType(1)  //	名		名
	Verb1        = WordType(2)  //	動Ⅰ		动Ⅰ
	Verb2        = WordType(3)  //	動Ⅱ		动Ⅱ
	Verb3        = WordType(4)  //	動Ⅲ		动Ⅲ
	AdjectiveI   = WordType(5)  //	イ形		い形
	AdjectiveNa  = WordType(6)  //	ナ形		な形
	Auxiliary    = WordType(8)  //	助詞		助词
	Pronoun      = WordType(9)  //	代名詞	代词
	Interjection = WordType(10) //	感		感叹词
	Adverb1      = WordType(11) //	副		副
	Adverb2      = WordType(12) //	副詞		副词
	Conjunction1 = WordType(13) //	連語		连语
	Conjunction2 = WordType(14) //	連体		连体
	Conjunction3 = WordType(15) //	接		接
	Conjunction4 = WordType(16) //	接詞		接词
	Conjunction5 = WordType(17) //	接尾		接尾
	BuildLang    = WordType(18) //			造语
	Quantifier   = WordType(19) //			助数
)

type BookVolume int8

const (
	NullVolume            = BookVolume(0)
	PrimaryVolumeOne      = BookVolume(1)
	PrimaryVolumeTwo      = BookVolume(2)
	IntermediateVolumeOne = BookVolume(3)
	IntermediateVolumeTwo = BookVolume(4)
)

type RelationType int8

//type WordRelation struct {
//	ID           int
//	WordA        Word `gorm:"primaryKey"`
//	WordB        Word `gorm:"primaryKey"`
//	RelationType RelationType
//}

type VocabularyCore struct {
	ID    uint   `gorm:"primaryKey;autoIncrement;"`
	Kana  string `gorm:"type:varchar(100);not null;comment:かな;index:idx_Kana_kanji,unique;"` // かな 假名
	Kanji string `gorm:"type:varchar(100);not null;comment:漢字;index:idx_Kana_kanji,unique;"` // 漢字 汉字
}

type Vocabulary struct {
	VocabularyCore
	Original bool `gorm:"default:true;not null;comment:是否是单词原形;"`
}

func NewVocabulary(vc VocabularyCore, original bool) *Vocabulary {
	return &Vocabulary{
		VocabularyCore: vc,
		Original:       original,
	}
}

type WordMeaning struct {
	ID          uint `gorm:"primaryKey;"`
	WordID      uint
	Major       bool       `gorm:"default:true;not null;comment:是否是最常见释义;"`
	WordType    WordType   `gorm:"not null;comment:名字类型;"`
	Meaning     string     `gorm:"type:varchar(2000);comment:释义;"`
	Description string     `gorm:"type:varchar(2000);comment:解释;"`
	BookVolume  BookVolume `gorm:"not null;comment:教材;"`
	UnitNo      int8       `gorm:"not null;comment:课程序号;"`
}

type Word struct {
	VocabularyCore
	WordMeanings []WordMeaning `gorm:"foreignKey:WordID;references:ID"`
	// WordRelations []*Word   `gorm:"many2many:word_relations"`
}
