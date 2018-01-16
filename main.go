package main

import (
	"fmt"
	"log"

	"github.com/fsufitch/sample-go-app/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{Name: "port, p", Usage: "port to listen on", Value: 8080, EnvVar: "PORT"},
	}
	app.ArgsUsage = "STATIC_DIR"
	app.Action = serve

	app.RunAndExitOnError()
}

func serve(ctx *cli.Context) error {
	port := ctx.Int("port")
	if port == 0 {
		fmt.Println("error: Specify port either via `-p` or via PORT env var")
		return ctx.App.Command("help").Run(ctx)
	}

	staticPath := ctx.Args().First()
	if staticPath == "" {
		fmt.Println("error: STATIC_DIR argument required")
		return ctx.App.Command("help").Run(ctx)
	}

	log.Printf("Sending port and static path to server module (%d, %s)", port, staticPath)
	return server.ListenAndServe(port, staticPath)
}
