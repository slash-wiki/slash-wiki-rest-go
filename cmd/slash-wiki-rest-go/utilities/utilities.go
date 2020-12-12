package utilities

import (
	Structures "github.com/they-them/slash-wiki-rest-go/cmd/slash-wiki-rest-go/structures" 
	"net/http"
)

func FormValuesToParameters(request *http.Request) Structures.Parameters {
	token := request.FormValue("token")
	teamId := request.FormValue("team_id")
	teamDomain := request.FormValue("team_domain")
	channelId := request.FormValue("channel_id")
	channelName := request.FormValue("channel_name")
	userId := request.FormValue("user_id")
	userName := request.FormValue("user_name")
	command := request.FormValue("command")
	text := request.FormValue("text")
	responseUrl := request.FormValue("response_url")
	triggerId := request.FormValue("trigger_id")
	apiAppId := request.FormValue("api_app_id")

	return Structures.Parameters{
		Token:       token,
		TeamId:      teamId,
		TeamDomain:  teamDomain,
		ChannelId:   channelId,
		ChannelName: channelName,
		UserId:      userId,
		UserName:    userName,
		Command:     command,
		Text:        text,
		ResponseUrl: responseUrl,
		TriggerId:   triggerId,
		ApiAppId:    apiAppId,
	}
}