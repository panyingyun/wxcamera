package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/panyingyun/wxcamera/httpserver"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func run(c *cli.Context) error {
	//start Restful API http server
	httpserver.Start(c.String("server-url"))
	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Infof("signal received signal %v", <-sigChan)
	log.Warn("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "weixin server"
	app.Usage = "weixin server"
	app.Copyright = "panyingyun@gmail.com"
	app.Version = "0.0.1"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server-url",
			Usage:  "server-url",
			Value:  "127.0.0.1:8902",
			EnvVar: "SERVER_URL",
		},
	}
	app.Run(os.Args)
}
