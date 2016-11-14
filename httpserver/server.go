package httpserver

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/go-macaron/cache"
	"gopkg.in/macaron.v1"
)

func Start(addr string) {
	//start http server
	go func() {
		startHttpServer(addr)
	}()
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{APPTOKEN, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

//For Server Get
func WXServerGet(ctx *macaron.Context) string {
	log.Info("WXServerGet begin")
	timestamp := ctx.Params("timestamp")
	nonce := ctx.Params("nonce")
	echostr := ctx.Params("echostr")
	signatureIn := ctx.Params("signature")

	signaturegen := makeSignature(timestamp, nonce)
	if signaturegen != signatureIn {
		log.Error("Validata sign fail!!!")
	}

	if echostr == "" {
		return "WXServerGet Here!!!"
	}
	return echostr
}

//For Server Post
func WXServerPost(ctx *macaron.Context) string {
	log.Info("WXServerPost begin")
	return "Welcome,Wx Server Here!!!"
}

func startHttpServer(addr string) {
	macaron.Env = macaron.PROD
	m := macaron.Classic()
	m.Use(cache.Cacher())
	m.Get("/", WXServerGet)
	m.Post("/", WXServerPost)
	m.RunAddr(addr)
}
