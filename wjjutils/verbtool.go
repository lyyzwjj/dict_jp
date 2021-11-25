package wjjutils

import (
	"github.com/bwmarrin/snowflake"
)

var NodeGen *snowflake.Node

func InitSnowflake() {
	var err error
	if NodeGen, err = snowflake.NewNode(1); err != nil {
		panic("InitSnowflake init failed!")
	}
}
