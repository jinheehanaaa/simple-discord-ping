// Responsible for sending the messages
package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jinheehanaaa/simple-discord-ping/config"
)

var BotID string
var goBot *discordgo.Session // Store Session into goBot

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token) // Create new bot session with token

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me") // Create the bot as a user

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID // Give Bot a ID

	goBot.AddHandler(messageHandler) // Read & Write messages to the channel

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Don't do anything
	if m.Author.ID == BotID {
		return
	}

	// Send message
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "pingping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pongpong")
	}

	if m.Content == "pingpingping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pongpongpong")
	}
}
