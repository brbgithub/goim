package connect

import (
	"github.com/brbgithub/goim/logic/db"
	"github.com/brbgithub/goim/public/imctx"
)

func Context() *imctx.Context{
	return imctx.NewContext(db.Factory.GetSession())
}
