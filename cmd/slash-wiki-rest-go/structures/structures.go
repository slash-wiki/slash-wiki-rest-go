package structures

type Parameters struct {
	// Payload fields from "Preparing your app to receive Commands" section in documentation:
	// https://api.slack.com/interactivity/slash-commands#app_command_handling.
	Token       string `schema:"token"`
	TeamId      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelId   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserId      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Command     string `schema:"command"`
	Text        string `schema:"text"`
	ResponseUrl string `schema:"response_url"`
	TriggerId   string `schema:"trigger_id"`
	ApiAppId    string `schema:"api_app_id"`
}

type Message struct {
	// Response type for returning Slack messages.
	ResponseType string `schema:"response_type"`
	Text         string `schema:"text"`
}