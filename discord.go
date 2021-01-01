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
	applicationCommandOptionTypeSubcommandGroup
	applicationCommandOptionTypeString
)

type applicationCommandOptionChoice struct {
	Name  string      `json:"name"`
	Value stringOrInt `json:"value"`
}

type stringOrInt struct {
	StrVal string
	IntVal string
}

type webhookRequest struct {
	Type webhookRequestType `json:"type"`
}

type webhookRequestType int

const (
	webhookRequestTypePing webhookRequestType = iota + 1
)

type webhookResponse struct {
	Type webhookRequestType `json:"type"`
}
