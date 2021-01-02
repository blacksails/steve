package steve

import (
	"encoding/hex"
	"net/http"
	"os"
	"sync"
)

var (
	cloudFunctionOnce    sync.Once
	cloudFunctionHandler http.HandlerFunc
)

func CloudFunction(w http.ResponseWriter, r *http.Request) {
	cloudFunctionOnce.Do(func() {

		s := New(
			AppID(os.Getenv("STEVE_APPLICATION_ID")),
			BotToken(os.Getenv("STEVE_BOT_TOKEN")),
			GuildID(os.Getenv("STEVE_GUILD_ID")),
		)

		apk, err := hex.DecodeString(os.Getenv("STEVE_APPLICATION_PUBKEY"))
		if err != nil {
			s.log.Error(err, "could not decode pubkey")
		}

		AppPubKey(apk)(s)

		if err := s.RegisterCommands(); err != nil {
			s.log.Error(err, "could not register commands")
		}

		cloudFunctionHandler = s.Handler()
	})

	cloudFunctionHandler(w, r)
}
