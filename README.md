# Bard <img src="https://www.gstatic.com/lamda/images/favicon_v1_150160cddff7f294ce30.svg" width="35px" />
Reverse engineered Google bard cli implemented in golang
[![Go Reference](https://pkg.go.dev/badge/github.com/ganeshk312/bard-go.svg)](https://pkg.go.dev/github.com/ganeshk312/bard-go)


## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release) (we test it with these).

### Getting Bard-go

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/ganeshk312/bard-go/bard"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `gin` package:

```sh
$ go get -u github.com/ganeshk312/bard-go/bard
```

## Authentication
Go to https://bard.google.com/

- F12 for console
- Copy the values
  - Session: Go to Application → Cookies → `__Secure-1PSID`. Copy the value of that cookie.
- Provide this token to the NewChatbot api

## Developer Documentation

The NewChatbot api returns a bard client of type [Chatbot](https://github.com/ganeshk312/bard-go/blob/eb076686d495/bard/Types.go#L5). 
Ask function takes message as a string and provides the bard [Response](https://github.com/ganeshk312/bard-go/blob/eb076686d495/bard/Types.go#LL16C4-L16C4) object. Bard session retained for the Chatbot client as the functionalities supported by Google bard.
Currently, Bard returns three drafts for each prompt, of type [Choice](https://github.com/ganeshk312/bard-go/blob/eb076686d495/bard/Types.go#L25) in the [Response](https://github.com/ganeshk312/bard-go/blob/eb076686d495/bard/Types.go#LL16C4-L16C4) object

### Sample Usage:

```go
chatbot, _ = bard.NewChatbot(sessionID)
response, err := chatbot.Ask(message)
```



