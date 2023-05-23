package bard

import "net/http"

type Chatbot struct {
	headers        http.Header
	reqID          int
	SNlM0e         string
	conversationID string
	responseID     string
	choiceID       string
	client         *http.Client
	sessionid      string
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
