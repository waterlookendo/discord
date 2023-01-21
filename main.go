package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("こんにちは世界")

	_, _ = discordgo.New("Bot " + "authentication token")
}
