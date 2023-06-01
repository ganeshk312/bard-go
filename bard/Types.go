package bard

import "net/http"

type Chatbot struct {
	ReqID          int
	SNlM0e         string
	ConversationID string
	ResponseID     string
	ChoiceID       string
	Client         *http.Client
	Sessionid      string
}

type Response struct {
	Content           string
	ConversationID    string
	ResponseID        string
	FactualityQueries []interface{}
	TextQuery         string
	Choices           []Choice
}

type Choice struct {
	ID      string
	Content string
}
