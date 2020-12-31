package main

import (
	"log"
	"net/http"
	"os"

	"github.com/blacksails/steve"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "steve",
		Usage: "discord minecraft guy",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "application-id",
				Usage:    "discord application id",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "bot-token",
				Usage:    "discord bot token",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "guild-id",
				Usage:    "discord guild id",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "run the server",
				Action: func(c *cli.Context) error {
					return http.ListenAndServe(":8080", newSteve(c).Handler())
				},
			},
			{
				Name:  "register-commands",
				Usage: "run the server",
				Action: func(c *cli.Context) error {
					return newSteve(c).RegisterCommands()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newSteve(c *cli.Context) *steve.Server {
	return steve.New(
		steve.AppID(c.String("application-id")),
		steve.BotToken(c.String("bot-token")),
		steve.GuildID(c.String("guild-id")),
	)
}
