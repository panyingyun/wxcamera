package httpserver

import (

	//log "github.com/Sirupsen/logrus"

	"github.com/go-macaron/cache"
	"gopkg.in/macaron.v1"
)

func Start(addr string) {
	//start http server
	go func() {
		startHttpServer(addr)
	}()
}

//For Server Test
func WXServer(ctx *macaron.Context) string {
	return "Welcome,Wx Server Here!!!"
}

func startHttpServer(addr string) {
	macaron.Env = macaron.PROD
	m := macaron.Classic()
	m.Use(cache.Cacher())
	m.Get("/", WXServer)
	m.RunAddr(addr)
}
