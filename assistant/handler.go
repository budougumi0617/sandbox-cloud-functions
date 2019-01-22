package assistant

import (
	"fmt"
	"net/http"
	"os"
)

// Hello is test func
func Hello(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("SLACK_URL")
	s := NewSlack(url, "Hi!!!!", "BOT", "")
	s.Send()
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "someone"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
