package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v6"
)

type envVars struct {
	BotToken	string 	`env:"BOT_TOKEN"`
}

func main() {
	envs := envVars{}
	if err := env.Parse(&envs); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", envs)

	_, _ = discordgo.New("Bot " + "authentication token")
}
