package main

import (
	"os"

	"github.com/go-chat-bot/bot/irc"
	_ "github.com/go-chat-bot/plugins/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	irc.Run(&irc.Config{
		Server:   os.Getenv("IRC_SERVER"),
		Channels: []string{os.Getenv("IRC_CHANNELS") + " " + os.Getenv("IRC_CHANNELS_PASSWORD")},
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != "",
	})
}
