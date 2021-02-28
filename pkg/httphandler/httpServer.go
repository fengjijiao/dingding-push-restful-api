package httphandler

import (
	"path"
    "net/http"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/conf"
)

func Run() error {
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, "/"), defaultHttpHandler)
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, conf.Config.SecurityPrefix, "send"), sendMessageHttpHandler)
    return http.ListenAndServe(conf.Config.HttpServerListen, nil)
}