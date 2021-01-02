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
	Member    string                            `json:"member,omitempty"`
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
