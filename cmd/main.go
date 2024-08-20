package main

import (
	"net/http"
	"os"

	"github.com/jpillora/overseer"
	overseerFetcher "github.com/jpillora/overseer/fetcher"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	_ "gl.eeo.im/fengye2419/ai-audio-service/docs"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/app/routers"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/setting"
)

// @title AI Audio Service API
// @version 1.0
// @description This is a sample server for AI Audio Service.
// @host localhost:8080
// @BasePath /
func main() {
	app := &cli.App{
		Name:  "ai-audio-service",
		Usage: "A web service with ai audio features",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILE`",
				Value: setting.ConfigFile,
			},
		},
		Action: func(c *cli.Context) error {
			// Load configuration
			configFile := c.String("config")
			setting.SetConfigFile(configFile)
			setting.Setup()
			// Start the server
			if setting.IsLocal() {
				localListen()
			} else {
				overseer.Run(overseer.Config{
					Program: prog,
					Address: setting.Overseer.HTTPAddress,
					Fetcher: &overseerFetcher.File{
						Path: setting.Overseer.FetcherFilePath,
					},
					Debug: setting.Overseer.DebugMode,
				})
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// prog is the main program overseer will manage
func prog(state overseer.State) {
	routers.GlobalInit()
	r := routers.NormalRoutes()
	if err := http.Serve(state.Listener, r); err != nil {
		logrus.Infof("http.Serve Listener closed ||overseer stateid: %s ||graceful shutdown: %v  ||err: %v", state.ID, <-state.GracefulShutdown, err)
	}
}

// localListen starts the server locally
func localListen() {
	routers.GlobalInit()
	r := routers.NormalRoutes()
	if err := http.ListenAndServe(setting.Overseer.HTTPAddress, r); err != nil {
		logrus.Fatalf("http.ListenAndServe: %v", err)
	}
}
