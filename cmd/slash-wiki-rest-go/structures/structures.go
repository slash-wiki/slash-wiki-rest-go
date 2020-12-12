package structures

type Parameters struct {
	// Payload fields from "Preparing your app to receive Commands" section in documentation:
	// https://api.slack.com/interactivity/slash-commands#app_command_handling.
	Token       string `json:"token"`
	TeamId      string `json:"team_id"`
	TeamDomain  string `json:"team_domain"`
	ChannelId   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Command     string `json:"command"`
	Text        string `json:"text"`
	ResponseUrl string `json:"response_url"`
	TriggerId   string `json:"trigger_id"`
	ApiAppId    string `json:"api_app_id"`
}

type Message struct {
	// Response type for returning Slack messages.
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}