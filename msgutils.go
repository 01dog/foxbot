package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("help", false, msgHelp).SetHelp("ok").Add()
}

func msgHelp(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	if len(msgList) == 2 {
		if v, ok := activeCommands[strings.ToLower(msgList[1])]; ok {
			v.helpCommand(s, m)
			return
		}

	}

	var commands []string
	for _, v := range activeCommands {
		if m.Author.ID == config.AdminID || !v.AdminOnly {
			commands = append(commands, "`"+v.Name+"`")
		}
	}

	em := NewEmbed().
		AddField("bot help", strings.Join(commands, ", ")).
		SetFooter("use .help [command] for a detailed info").MessageEmbed
	s.ChannelMessageSendEmbed(m.ChannelID, em)
}

func (c Command) helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	em := NewEmbed().
		SetColor(0).
		AddField(c.Name, c.Help).MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}
