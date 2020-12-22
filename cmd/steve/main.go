package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "steve",
		Usage: "discord minecraft guy",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "run the server",
				Action: func(c *cli.Context) error {
					return runServer()
				},
			},
			{
				Name:  "register-commands",
				Usage: "run the server",
				Action: func(c *cli.Context) error {
					return registerCommands(
						c.String("application-id"),
						c.String("bot-token"),
						c.String("guild-id"),
					)
				},
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func registerCommands(appID, botToken, guildID string) error {
	cmd := ApplicationCommand{
		Name:        "steve",
		Description: "control minecraft server",
		Options: []ApplicationCommandOption{
			{
				Name:        "whitelist",
				Description: "whitelist a minecraft username",
				Type:        ApplicationCommandOptionTypeSubcommand,
			},
		},
	}

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(cmd); err != nil {
		return errors.Wrap(err, "could not encode application command")
	}
	fmt.Println(buf.String())

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://discord.com/api/v8/applications/%s/guilds/%s/commands", appID, guildID),
		&buf,
	)
	if err != nil {
		return errors.Wrap(err, "could not create http request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", botToken))
	req.Header.Set("Content-Type", "application/json")

	hc := http.Client{}
	res, err := hc.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not send http request")
	}

	resB := map[string]interface{}{}
	json.NewDecoder(res.Body).Decode(&resB)
	fmt.Println(res.StatusCode)
	fmt.Printf("%+v\n", resB)
	return nil
}

type ApplicationCommand struct {
	ID            string                     `json:"id,omitempty"`
	ApplicationID string                     `json:"application_id,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Description   string                     `json:"description,omitempty"`
	Options       []ApplicationCommandOption `json:"options,omitempty"`
}

type ApplicationCommandOption struct {
	Type        ApplicationCommandOptionType     `json:"type"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Default     *bool                            `json:"default,omitempty"`
	Required    *bool                            `json:"required,omitempty"`
	Choices     []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options     []ApplicationCommandOption       `json:"options,omitempty"`
}

type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubcommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubcommandGroup
	ApplicationCommandOptionTypeString
)

type ApplicationCommandOptionChoice struct {
	Name  string      `json:"name"`
	Value StringOrInt `json:"value"`
}

type StringOrInt struct {
	StrVal string
	IntVal string
}

func runServer() error {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	return http.ListenAndServe(":8080", r)
}
