package service

import (
	"database/sql"
	"github.com/brbgithub/goim/logic/dao"
	"github.com/brbgithub/goim/public/imctx"
	"github.com/brbgithub/goim/public/imerror"
)

type authService struct{}

var AuthService = new(authService)

func (*authService) Auth(ctx *imctx.Context, deviceId int64, token string) (int64, error) {
	device, err := dao.DeviceDao.Get(ctx, deviceId)
	if err == sql.ErrNoRows {
		return 0, imerror.LErrDeviceIdOrToken
	}
	if err != nil {
		return 0, err
	}

	if token != device.Token {
		return 0, imerror.LErrDeviceIdOrToken
	}

	return device.UserId, nil
}
