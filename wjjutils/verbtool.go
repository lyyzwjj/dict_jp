package wjjutils

import (
	"github.com/bwmarrin/snowflake"
)

var NodeGen *snowflake.Node

func init() {
	var err error
	if NodeGen, err = snowflake.NewNode(1); err != nil {
		panic("InitSnowflake init failed!")
	}
}

func GenID() uint64 {
	return uint64(NodeGen.Generate())
}
