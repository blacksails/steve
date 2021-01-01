package steve

import (
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
		cloudFunctionHandler = s.Handler()
	})

	cloudFunctionHandler(w, r)
}
