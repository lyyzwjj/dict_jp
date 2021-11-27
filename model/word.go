package model

type WordType uint8

const (
	Sentence     = WordType(0)  //	句子		句子
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

type BookVolume uint8

const (
	NullVolume            = BookVolume(0)
	PrimaryVolumeOne      = BookVolume(1)
	PrimaryVolumeTwo      = BookVolume(2)
	IntermediateVolumeOne = BookVolume(3)
	IntermediateVolumeTwo = BookVolume(4)
)

type PitchAccent int

const (
	PitchAccentLvNil = PitchAccent(-1)
	PitchAccentLv0   = PitchAccent(0 << 1)
	PitchAccentLv1   = PitchAccent(1 << 1)
	PitchAccentLv2   = PitchAccent(2 << 1)
	PitchAccentLv3   = PitchAccent(3 << 1)
	PitchAccentLv4   = PitchAccent(4 << 1)
	PitchAccentLv5   = PitchAccent(5 << 1)
	PitchAccentLv6   = PitchAccent(6 << 1)
	PitchAccentLv7   = PitchAccent(7 << 1)
	PitchAccentLv8   = PitchAccent(8 << 1)
	PitchAccentLv9   = PitchAccent(9 << 1)
	PitchAccentLv10  = PitchAccent(10 << 1)
	PitchAccentLv11  = PitchAccent(11 << 1)
)

type RelationType uint8

//type WordRelation struct {
//	ID           int
//	WordA        Word `gorm:"primaryKey"`
//	WordB        Word `gorm:"primaryKey"`
//	RelationType RelationType
//}

type VocabularyCore struct {
	ID     uint64 `gorm:"type:bigint;primaryKey;autoIncrement:false;"`
	Kana   string `gorm:"type:varchar(255);not null;comment:かな;" // index:idx_Kana_kanji,unique;"` // かな 假名
	Kanji  string `gorm:"type:varchar(255);not null;comment:漢字;" // index:idx_Kana_kanji,unique;"` // 漢字 汉字
	Romaji string `gorm:"type:varchar(255);not null;comment:ロマ字;"`
}

type Vocabulary struct {
	VocabularyCore
	Original bool `gorm:"default:true;not null;comment:是否是单词原形;"`
}

func NewVocabulary(vocabularyCore VocabularyCore, original bool) *Vocabulary {
	return &Vocabulary{
		VocabularyCore: vocabularyCore,
		Original:       original,
	}
}

type WordMeaning struct {
	ID          uint64     `gorm:"type:bigint;primaryKey;autoIncrement:false;"`
	WordID      uint       `gorm:"not null;comment:关联Word表主键;"`
	WordType    WordType   `gorm:"not null;comment:名字类型;"`
	Meaning     string     `gorm:"type:varchar(2000);comment:释义;"`
	Description string     `gorm:"type:varchar(2000);comment:解释;"`
	BookVolume  BookVolume `gorm:"not null;default:0;comment:教材;"`
	UnitNo      uint8      `gorm:"not null;default:0;comment:课程序号;"`
}

type Word struct {
	VocabularyCore
	WordType          WordType      `gorm:"not null;comment:名字类型;"`
	Meaning           string        `gorm:"type:varchar(2000);comment:释义;"`
	Description       string        `gorm:"type:varchar(2000);comment:解释;"`
	Book              string        `gorm:"type:varchar(20);not null;comment:教材;"`
	UnitNo            uint8         `gorm:"not null;default:0;comment:课程序号;"`
	Masu              string        `gorm:"type:varchar(255);not null;comment:ます形,动词专有;"`
	PitchAccentFirst  PitchAccent   `gorm:"type:tinyint;not null;default:-1;comment:音调;"`
	PitchAccentSecond PitchAccent   `gorm:"type:tinyint;not null;default:-1;comment:音调;"`
	WordMeanings      []WordMeaning `gorm:"foreignKey:WordID;references:ID"`
	QueryCount        uint          `gorm:"not null;default:0;comment:查询次数;"`
	// WordRelations []*Word   `gorm:"many2many:word_relations"`
}
