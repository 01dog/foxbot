package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("lizard", false, lizard).Add()
}

func lizard(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	msgList = append(msgList, "reptiles")
	RedditBot(s, m, msgList)
}
