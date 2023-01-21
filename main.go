package main

import (
	"log"
	"os"
	"os/signal"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v6"
)

type envVars struct {
	BotToken	string 	`env:"BOT_TOKEN"`
	GuildID         string  `env:"GUILD_ID"`
}

func main() {
	envs := envVars{}
	if err := env.Parse(&envs); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", envs)

	dg, err := discordgo.New("Bot " + envs.BotToken)
	if err != nil {
		panic(err.Error())
	}

	if err != nil { panic(err.Error()) }

	// dg.Identify.Intents = discordgo.
	dg.AddHandler(SlashCommandHandler);

	/* grab the bot user's id */
	botUser, err := dg.User("@me")
	if err != nil { log.Panic(err.Error()) }
	_ = botUser.ID

	err = dg.Open()
	if err != nil { log.Panic(err.Error()) }

	for _, c := range SlashCommands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, envs.GuildID, c)
		if err != nil {
			log.Printf("Failed to create slash command: %v", c.Name)
		}
	}

	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

}

var SlashCommands = []*discordgo.ApplicationCommand{
	{
		Name: "hello",
		Type: discordgo.ChatApplicationCommand,
		Description: "hello world",
	},
}

func SlashCommandHandler(s *discordgo.Session, m *discordgo.InteractionCreate) {
	if handler, ok := SlashCommandHandlers[m.ApplicationCommandData().Name]; ok {
		handler(s, m)
	}
}

var SlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "hello world slash command",
				            
			},
			        
		})
	},
}
