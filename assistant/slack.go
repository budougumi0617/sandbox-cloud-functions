package assistant

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// Slack sends message.
type Slack struct {
	params params
	url    string
}

type params struct {
	Text      string `json:"text"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
	Channel   string `json:"channel,omitempty"`
}

// NewSlack returns FIXME
func NewSlack(url string, text string, username string, channel string) *Slack {

	e := ":ghost:"
	if len(channel) == 0 {
		channel = "#general"
	}

	p := params{
		Text:      text,
		Username:  username,
		IconEmoji: e,
		//IconURL:   iconURL,
		Channel: channel}

	return &Slack{
		params: p,
		url:    url,
	}
}

// Send sends message.
func (s *Slack) Send() {

	params, _ := json.Marshal(s.params)

	resp, err := http.PostForm(
		s.url,
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
}
