package stevefunc

import (
	"net/http"
	"os"
	"sync"

	"github.com/blacksails/steve"
)

var (
	fOnce    sync.Once
	fHandler http.HandlerFunc
)

func Function(w http.ResponseWriter, r *http.Request) {
	fOnce.Do(func() {
		s := steve.New(
			steve.AppID(os.Getenv("STEVE_APPLICATION_ID")),
			steve.BotToken(os.Getenv("STEVE_BOT_TOKEN")),
			steve.GuildID(os.Getenv("STEVE_GUILD_ID")),
		)
		fHandler = s.Handler()
	})

	fHandler(w, r)
}
