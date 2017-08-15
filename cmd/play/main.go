package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gernest/utron"
	"github.com/urfave/cli"
	"github.com/utronframework/play"
)

func main() {
	a := cli.NewApp()
	a.Name = "play"
	a.Description = `
	Pure Go frontend for the go playground using utron and gu
	`
	a.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "starts the server",
			Action: serve,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port",
					Usage: "port to bind the server",
					Value: 8000,
				},
			},
		},
	}
	if err := a.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func serve(ctx *cli.Context) error {
	port := ctx.Int("port")
	a := utron.NewApp()
	a.Router.Add(play.New)
	log.Printf("starting server on port %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a)
}
