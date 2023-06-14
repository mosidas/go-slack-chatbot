package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const (
	postMessageUrl = "https://slack.com/api/chat.postMessage"
	uploadFileUrl  = "https://slack.com/api/files.upload"
)

type SlackClient struct {
	token   string
	channel string
}

func NewSlackClient(token string, channel string) *SlackClient {
	return &SlackClient{token: token, channel: channel}
}

func (s *SlackClient) PostMessageAsync(text string) (*ResponsePostMessage, error) {
	data := url.Values{
		"token":   {s.token},
		"channel": {s.channel},
		"text":    {text},
	}
	return s.postMessage(postMessageUrl, data)
}

func (s *SlackClient) ReplyMessageAsync(text string, ts string) (*ResponsePostMessage, error) {
	data := url.Values{
		"token":     {s.token},
		"channel":   {s.channel},
		"text":      {text},
		"thread_ts": {ts},
	}
	return s.postMessage(postMessageUrl, data)
}

func (s *SlackClient) postMessage(url string, data url.Values) (*ResponsePostMessage, error) {
	response, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var respJson ResponsePostMessage
	err = json.NewDecoder(response.Body).Decode(&respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}

func (s *SlackClient) UploadFileAsync(filePath string, text string) (*ResponseFileUpload, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	io.Copy(part, file)

	writer.WriteField("token", s.token)
	writer.WriteField("channels", s.channel)
	writer.WriteField("filename", filepath.Base(filePath))
	writer.WriteField("initial_comment", text)

	req, err := http.NewRequest("POST", uploadFileUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respJson ResponseFileUpload
	err = json.NewDecoder(resp.Body).Decode(&respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}
