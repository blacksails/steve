package steve

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
