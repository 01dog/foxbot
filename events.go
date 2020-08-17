package main

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// EvalMessage will parse all messages that get sent and determine if someone issues a command
func EvalMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore the message if its from the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// regex to match the N word
	// this is because of you, Gust. . .
	r, _ := regexp.Compile("(?i)n[ei1]+(g|6)+(ger|ga|a|6er|g3r|64|4|63r|e|r|ge|gr)[sz]*")
	message := strings.Split(m.Content, " ")
	joinedMessage := strings.Join(message, "")

	if match := r.Match([]byte(joinedMessage)); match {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
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
