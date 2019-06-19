package service

import (
	"fmt"
	"github.com/brbgithub/goim/logic/db"
	"github.com/brbgithub/goim/public/imctx"
	"testing"
)

var ctx = imctx.NewContext(db.Factory.GetSession())

func TestFriendService_Add(t *testing.T) {
	err := FriendService.Add(ctx, 1,2,"brbgithub")
	if err != nil {
		fmt.Println(err)
	}
}

func TestFriendService_Delete(t *testing.T) {
	err := FriendService.Delete(ctx, 1, 2)
	if err != nil {
		fmt.Println(err)
	}
}
