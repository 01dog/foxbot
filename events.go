package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// EvalMessage will parse all messages that get sent and determine if someone issues a command
func EvalMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore the message if its from the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// ignore the message if it doesnt have the command prefix
	if !strings.HasPrefix(m.Content, config.Prefix) {
		return
	}

	ParseCommand(s, m, func() string {
		if strings.HasPrefix(m.Content, config.Prefix) {
			return strings.TrimPrefix(m.Content, config.Prefix)
		}
		return strings.TrimPrefix(m.Content, config.Prefix)
	}())
}
