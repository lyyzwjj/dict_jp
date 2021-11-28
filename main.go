package main

import (
	"github.com/lyyzwjj/dict_jp/dao"
	"github.com/lyyzwjj/dict_jp/model"
)

func main() {
	//fmt.Println(strings.Compare("黒", "白"))
	//fmt.Println(strings.Compare("くろ", "しろ"))
	//fmt.Println(strings.Compare("kuro", "siro"))
	//fmt.Println(strings.Compare("ううん", "うえ"))
	//fmt.Println(strings.Compare("ううん", "うえます"))
	//fmt.Println(strings.Compare("ううん", "あ"))
	//fmt.Println(strings.Compare("ううん", "じゅんび"))
	//fmt.Println(strings.Compare("ううん", "ジョジョと"))
	//wjjutils.A()
	// dao.Repo.AutoMigrate(&model.Vocabulary{})
	//var err = dao.Repo.SetupJoinTable(&model.Word{}, "WordRelations", &model.WordRelation{})
	//if err != nil {
	//	panic("join table setup failed!")
	//}
	dao.Repo.AutoMigrate(&model.Vocabulary{}, &model.Word{}, &model.WordBook{})
}
