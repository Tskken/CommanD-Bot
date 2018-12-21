package core

// TODO: Figure out permission shit...

type BotPermission map[string]Permissions

var BotPermissions = make(BotPermission)

type Permissions struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
