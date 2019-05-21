package imctx

import "github.com/brbgithub/goim/public/session"

type Context struct {
	Session *session.Session
}

func NewContext(Session *session.Session) *Context {
	return &Context{Session:Session}
}
