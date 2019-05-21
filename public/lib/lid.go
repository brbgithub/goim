package lib

import (
	"database/sql"
	"fmt"
	"github.com/brbgithub/goim/conf"
	"github.com/brbgithub/goim/public/lib/lid"
	"github.com/brbgithub/goim/public/logger"
)

var Lid *lid.Lid

func init() {
	db,err := sql.Open("mysql",conf.MySQL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	Lid,err = lid.NewLid(db,"message_id", 1000)
	if err != nil{
		logger.Sugar.Error(err)
		panic(err)
	}
}
