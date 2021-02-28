package conf

import (
	"github.com/fengjijiao/dingding-push-restful-api/pkg/logio"
)

type ConfInfo struct {
	WorkDir string `yaml:"work-dir"`
	DingdingAppKey string `yaml:"app-key"`
	DingdingAppSecret string `yaml:"app-secret"`
	DingdingAgentId	int	`yaml:"agent-id"`
	HttpServerListen string `yaml:"http-server-listen"`
	BaseUrlPath string `yaml:"base-url-path"`
	SecurityPrefix string `yaml:"security-prefix"`
	// DingdingWebHookToken string `yaml:"webhook-token"`
	// DingdingWebHookEnCodingAesKey string	`yaml:"webhook-encoding-aeskey"`
}

func (ci *ConfInfo) setDefaults() {
	if ci.WorkDir == "" {
		ci.WorkDir = "./"
	}
	if ci.DingdingAppKey == "" {
		logio.Logger.Fatal("[setDefaults]: DingdingAppKey can not be empty!")
	}
	if ci.DingdingAppSecret == "" {
		logio.Logger.Fatal("[setDefaults]: DingdingAppSecret can not be empty!")
	}
	if ci.HttpServerListen == "" {
		ci.HttpServerListen = ":9465"
	}
}