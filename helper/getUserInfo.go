package helper

import (
	"context"
	"ekak_kab_sleman/model/web"
)

// see contextkey.go
func GetUserInfo(ctx context.Context) web.JWTClaim {
	if ctx == nil {
		return web.JWTClaim{}
	}
	val := ctx.Value(UserInfoKey)
	if val == nil {
		return web.JWTClaim{}
	}

	claims, ok := val.(web.JWTClaim)
	if !ok {
		return web.JWTClaim{}
	}

	return claims
}
