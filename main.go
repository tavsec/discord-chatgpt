package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var authToken = os.Getenv("DISCORD_AUTH_TOKEN")
var chatAuthToken = os.Getenv("OPENAI_AUTH_TOKEN")

func init() {
	if authToken == "" {
		panic("You must set authentication token by setting DISCORD_AUTH_TOKEN environmental variable")
	}
	if chatAuthToken == "" {
		panic("You must set authentication token for OPEN AI by setting OPENAI_AUTH_TOKEN environmental variable")
	}
}

func main() {
	s, err := discordgo.New("Bot " + authToken)
	if err != nil {
		panic(err)
	}

	chatgpt := Client{
		AuthToken: chatAuthToken,
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == s.State.User.ID {
			return
		}

		log.Println("Generating answer for: " + m.Content)
		answer, err := chatgpt.Ask(m.Content)
		if err != nil {
			log.Fatalln(err)
		}
		s.ChannelMessageSend(m.ChannelID, answer.Choices[0].Message.Content)

	})
	s.Identify.Intents = discordgo.IntentsGuildMessages

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
