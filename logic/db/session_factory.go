package db

import (
	"github.com/brbgithub/goim/conf"
	"github.com/brbgithub/goim/public/session"
)

var Factory *session.SessionFactory

func init() {
	var err error
	Factory, err = session.NewSessionFactory("mysql",conf.MySQL)
	if err != nil{
		panic(err)
	}
}
