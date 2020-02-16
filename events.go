package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// EvalMessage will parse all messages that get sent and determine if someone issues a command
func EvalMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
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
