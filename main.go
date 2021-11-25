package main

import "github.com/lyyzwjj/dict_jp/dao"

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
	dao.InitMySQLDB()
}
