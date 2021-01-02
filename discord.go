package steve

type applicationCommand struct {
	ID            string                     `json:"id,omitempty"`
	ApplicationID string                     `json:"application_id,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Description   string                     `json:"description,omitempty"`
	Options       []applicationCommandOption `json:"options,omitempty"`
}

type applicationCommandOption struct {
	Type        applicationCommandOptionType     `json:"type"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Default     *bool                            `json:"default,omitempty"`
	Required    *bool                            `json:"required,omitempty"`
	Choices     []applicationCommandOptionChoice `json:"choices,omitempty"`
	Options     []applicationCommandOption       `json:"options,omitempty"`
}

type applicationCommandOptionType int

const (
	applicationCommandOptionTypeSubcommand applicationCommandOptionType = iota + 1
)

type applicationCommandOptionChoice struct {
	Name  string      `json:"name"`
	Value stringOrInt `json:"value"`
}

type stringOrInt struct {
	StrVal string
	IntVal string
}

type interaction struct {
	ID        string                            `json:"id,omitempty"`
	Type      interactionType                   `json:"type,omitempty"`
	Data      applicationCommandInteractionData `json:"data,omitempty"`
	GuildID   string                            `json:"guild_id,omitempty"`
	ChannelID string                            `json:"channel_id,omitempty"`
	Member    guildMember                       `json:"member,omitempty"`
	Token     string                            `json:"token,omitempty"`
	Version   int                               `json:"version,omitempty"`
}

type interactionType int

const (
	interactionTypePing interactionType = iota + 1
)

type applicationCommandInteractionData struct {
	ID      string                                    `json:"id,omitempty"`
	Name    string                                    `json:"name,omitempty"`
	Options []applicationCommandInteractionDataOption `json:"options,omitempty"`
}

type applicationCommandInteractionDataOption struct {
	Name    string                                    `json:"name,omitempty"`
	Value   string                                    `json:"value,omitempty"`
	Options []applicationCommandInteractionDataOption `json:"options,omitempty"`
}

type guildMember struct {
	User         user     `json:"user,omitempty"`
	Nick         string   `json:"nick,omitempty"`
	Roles        []string `json:"roles,omitempty"`
	JoinedAt     string   `json:"joined_at,omitempty"`
	PremiumSince string   `json:"premium_since,omitempty"`
	Deaf         bool     `json:"deaf,omitempty"`
	Mute         bool     `json:"mute,omitempty"`
	Pending      bool     `json:"pending,omitempty"`
}

type user struct {
	ID            string `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Discriminator string `json:"discriminator,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Bot           string `json:"bot,omitempty"`
	System        bool   `json:"system,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Email         string `json:"email,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	PremiumType   int    `json:"premium_type,omitempty"`
	PublicFlags   int    `json:"public_flags,omitempty"`
}
