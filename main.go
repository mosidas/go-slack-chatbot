package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./settings.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	token := viper.GetString("token")
	channel := viper.GetString("channel")

	if token == "" {
		log.Fatalf("Token is not provided in settings file")
	}
	if channel == "" {
		log.Fatalf("Channel is not provided in settings file")
	}

	client := NewSlackClient(token, channel)

	text := `
	<!channel>
	<@{user-id}>
	:slack::slack::slack::slack:
	Hello World!!!!あいうえお漢字①
	:slack::slack::slack::slack:
	:smile:
	:fire:
	:lion_face:
	`

	response, err := client.PostMessageAsync(text)
	if err != nil || response == nil {
		log.Fatalf("Error while posting message: %s", err)
	}
	if !response.Ok {
		log.Fatalf("Failed to post message: %s", response.Error)
	}
	log.Println(response.Ok)

	response, err = client.ReplyMessageAsync("Reply", response.Ts)
	if err != nil || response == nil {
		log.Fatalf("Error while replaying message: %s", err)
	}
	if !response.Ok {
		log.Fatalf("Failed to replay message: %s", response.Error)
	}
	log.Println(response.Ok)

	fileUploadResponse, err := client.UploadFileAsync("C:/tmp0/genba-neko.png", text)
	if err != nil || fileUploadResponse == nil {
		log.Fatalf("Error while uploading file: %s", err)
	}
	if !fileUploadResponse.Ok {
		log.Fatalf("Failed to upload file: %s", "fileUploadResponse.Error")
	}
	log.Println(fileUploadResponse.Ok)
}
