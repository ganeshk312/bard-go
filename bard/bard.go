package bard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const HOST = "bard.google.com"
const Origin_url = "https://" + HOST
const BASE_URL = "https://" + HOST + "/"
const ASK_URL = BASE_URL + "_/BardChatUi/data/assistant.lamda.BardFrontendService/StreamGenerate"

func NewChatbot(sessionID string) *Chatbot {
	headers := http.Header{
		"Host":          []string{HOST},
		"X-Same-Domain": []string{"1"},
		"User-Agent":    []string{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"},
		"Content-Type":  []string{"application/x-www-form-urlencoded;charset=UTF-8"},
		"Origin":        []string{Origin_url},
		"Referer":       []string{BASE_URL},
	}

	client := &http.Client{
		Timeout: 100 * time.Second,
	}
	client.Jar, _ = cookiejar.New(nil)
	chatBot := &Chatbot{
		headers:        headers,
		reqID:          rand.Intn(100000),
		SNlM0e:         "",
		conversationID: "",
		responseID:     "",
		choiceID:       "",
		client:         client,
		sessionid:      sessionID,
	}
	chatBot.setCookie(sessionID)
	chatBot.SNlM0e, _ = chatBot.getSNlM0e()
	return chatBot
}

func (chatBot *Chatbot) setCookie(sessionID string) {
	url, _ := url.Parse(BASE_URL)
	cookie := &http.Cookie{Name: "__Secure-1PSID", Value: sessionID}
	chatBot.client.Jar.SetCookies(url, []*http.Cookie{cookie})
}

func (c *Chatbot) getSNlM0e() (string, error) {
	resp, err := c.client.Get(BASE_URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not get google bard")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read the response body: ", err)
		return "", err
	}

	// Convert the response body to a string
	bodyString := string(body)

	re := regexp.MustCompile(`SNlM0e":"(.*?)"`)
	match := re.FindStringSubmatch(bodyString)
	if len(match) < 2 {
		return "", fmt.Errorf("SNlM0e not found")
	}
	// return "AFuTz6sM_HRJNfl9gxy7VknnKPuC:1684396254403", nil
	return match[1], nil
}

func (c *Chatbot) Ask(message string) (*Response, error) {
	params := url.Values{
		"bl":     {"boq_assistant-bard-web-server_20230514.20_p0"},
		"_reqid": {fmt.Sprintf("%d", c.reqID)},
		"rt":     {"c"},
	}
	data := url.Values{
		"f.req": []string{fmt.Sprintf(`[null, "[[\"%s\"], null, [\"%s\", \"%s\", \"%s\"]]"]`, message, c.conversationID, c.responseID, c.choiceID)},
		"at":    []string{c.SNlM0e},
	}
	log.Println(data)
	url := ASK_URL + "?" + params.Encode()
	req, _ := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header = c.headers
	c.setCookie(c.sessionid)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)
	respLines := strings.Split(buf.String(), "\n")
	respJSON := respLines[3]
	var jsonChatData []interface{}
	err = json.Unmarshal(json.RawMessage(respJSON), &jsonChatData)
	if err != nil {
		return nil, err
	}
	log.Println(jsonChatData[0])
	jsonChatData = jsonChatData[0].([]interface{})
	log.Println(jsonChatData[2])
	jsonChat := jsonChatData[2].(string)
	log.Println(jsonChat)
	err = json.Unmarshal(json.RawMessage(jsonChat), &jsonChatData)
	if err != nil {
		return nil, err
	}
	log.Println(jsonChatData)
	var choices []Choice
	for _, item := range jsonChatData[4].([]interface{}) {
		choices = append(choices, Choice{ID: item.([]interface{})[0].(string), Content: item.([]interface{})[1].([]interface{})[0].(string)})
	}
	results := &Response{
		Content:           jsonChatData[0].([]interface{})[0].(string),
		ConversationID:    jsonChatData[1].([]interface{})[0].(string),
		ResponseID:        jsonChatData[1].([]interface{})[1].(string),
		FactualityQueries: jsonChatData[3].([]interface{}),
		TextQuery:         jsonChatData[2].([]interface{})[0].([]interface{})[0].(string),
		Choices:           choices,
	}
	c.conversationID = results.ConversationID
	c.responseID = results.ResponseID
	c.choiceID = results.Choices[0].ID
	c.reqID += 100000
	return results, nil

}
