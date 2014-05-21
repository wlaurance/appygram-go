package appygram

type AppygramMessage struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Message  string `json:"message,omitempty"`
	Platform string `json:"platform,omitempty"`
	Software string `json:"software,omitempty"`
	Summary  string `json:"summary,omitempty"`
	Topic    string `json:"topic,omitempty"`
	AppJson  string `json:"app_json,omitempty"`
}

type AppygramTrace struct {
	AppygramMessage
	Class     string
	Backtrace [][]string
}

type AppygramResult struct {
	OK bool
}

type AppygramTopic struct {
	Name string
	Id   string
}

type AppygramTopics struct {
	AppygramResult
	Topics []AppygramTopic
}

type AppygramMessageWithKey struct {
	ApiKey string `json:"api_key",omitempty`
	AppygramMessage
}
