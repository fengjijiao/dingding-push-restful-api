package commonio

import (
	"github.com/fengjijiao/dingding-push-restful-api/pkg/conf"
)

func GetTokenAuto() (*AccessTokenInfo, error) {
	return GetToken(conf.Config.DingdingAppKey, conf.Config.DingdingAppSecret)
}