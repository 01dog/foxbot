package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("help", false, msgHelp).SetHelp("ok").Add()
}

func msgHelp(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	// check if help was given a specific (and valid) command
	if len(msgList) == 2 {
		if v, ok := activeCommands[strings.ToLower(msgList[1])]; ok {
			v.helpCommand(s, m)
			return
		}
	}

	// gets a list of active commands
	var commands []string
	for _, v := range activeCommands {
		if IsInArray(config.AdminID, m.Author.ID) || !v.AdminOnly {
			commands = append(commands, "`"+v.Name+"`")
		}
	}

	em := NewEmbed().
		AddField("bot help", strings.Join(commands, ", ")).
		SetFooter("use >help [command] for a detailed info").MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}

// if help is given a specific command, this gets the Help field and sends it in an embed
func (c Command) helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	em := NewEmbed().
		SetColor(0).
		AddField(c.Name, c.Help).MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}
