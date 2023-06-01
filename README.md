# Bard <img src="https://www.gstatic.com/lamda/images/favicon_v1_150160cddff7f294ce30.svg" width="35px" />
Reverse engineered Google bard cli implemented in golang
[![Go Reference](https://pkg.go.dev/badge/github.com/ganeshk312/bard-go.svg)](https://pkg.go.dev/github.com/ganeshk312/bard-go)

## Authentication
Go to https://bard.google.com/

- F12 for console
- Copy the values
  - Session: Go to Application → Cookies → `__Secure-1PSID`. Copy the value of that cookie.

## Developer Documentation


```go
chatbot, _ = bard.NewChatbot(sessionID)
response, err := chatbot.Ask(message)
```



Got the idea from [Antonio Cheong](https://github.com/acheong08/Bard)

