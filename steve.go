package steve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/pkg/errors"
)

type Server struct {
	log      logr.Logger
	appID    string
	botToken string
	guildID  string
}

type Option func(*Server)

func Logger(logger logr.Logger) Option {
	return func(s *Server) {
		s.log = logger
	}
}

func AppID(id string) Option {
	return func(s *Server) {
		s.appID = id
	}
}

func BotToken(token string) Option {
	return func(s *Server) {
		s.botToken = token
	}
}

func GuildID(id string) Option {
	return func(s *Server) {
		s.guildID = id
	}
}

func New(opts ...Option) *Server {
	var s Server
	for _, opt := range opts {
		opt(&s)
	}
	if s.log == nil {
		s.log = stdr.New(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))
	}
	return &s
}

func (s *Server) RegisterCommands() error {
	cmd := ApplicationCommand{
		Name:        "steve",
		Description: "control minecraft server",
		Options: []ApplicationCommandOption{
			{
				Name:        "whitelist",
				Description: "whitelist a minecraft username",
				Type:        ApplicationCommandOptionTypeSubcommand,
			},
			{
				Name:        "say",
				Description: "say something in the minecraft chat",
				Type:        ApplicationCommandOptionTypeSubcommand,
			},
		},
	}

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(cmd); err != nil {
		return errors.Wrap(err, "could not encode application command")
	}
	s.log.WithValues(
		"body", buf.String(),
	).Info("request")

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://discord.com/api/v8/applications/%s/guilds/%s/commands", s.appID, s.guildID),
		&buf,
	)
	if err != nil {
		return errors.Wrap(err, "could not create http request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", s.botToken))
	req.Header.Set("Content-Type", "application/json")

	hc := http.Client{}
	res, err := hc.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not send http request")
	}

	resB := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(&resB); err != nil {
		return errors.Wrap(err, "could not decode response")
	}
	s.log.WithValues(
		"status", res.StatusCode,
		"body", fmt.Sprintf("%+v\n", resB),
	).Info("got response")

	return nil
}
