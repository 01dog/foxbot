package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("lizard", false, lizard).Add()
}

func lizard(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	msgList = msgList[0:1]
	msgList = append(msgList, "reptiles")
	RedditBot(s, m, msgList)
}
