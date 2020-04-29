package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var activeCommands = make(map[string]Command)

//Command ...
type Command struct {
	Name      string
	Help      string
	AdminOnly bool
	Exec      func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

//ParseCommand checks for valid commands and permissions
func ParseCommand(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
	msgList := strings.Fields(message)
	if len(msgList) == 0 {
		return
	}

	commandName := strings.ToLower(func() string {
		if strings.HasPrefix(message, " ") {
			return " " + msgList[0]
		}
		return msgList[0]
	}())

	if command, ok := activeCommands[commandName]; ok && commandName == strings.ToLower(command.Name) {
		isAdmin := InArray(config.AdminID, m.Author.ID)
		if !command.AdminOnly || isAdmin {
			command.Exec(s, m, msgList)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "you do not have permission to use this command!")
		return
	}
}

//Add will add our command to the active command map
func (c Command) Add() Command {
	activeCommands[strings.ToLower(c.Name)] = c
	return c
}

//NewCommand will create our new Command
func NewCommand(name string, needAdmin bool, f func(*discordgo.Session, *discordgo.MessageCreate, []string)) Command {
	return Command{
		Name:      name,
		AdminOnly: needAdmin,
		Exec:      f,
	}
}

//Alias ...
func (c Command) Alias(a string) Command {
	activeCommands[strings.ToLower(a)] = c
	return c
}

//SetHelp ...
func (c Command) SetHelp(help string) Command {
	c.Help = help
	return c
}

//ReqAdmin ...
func (c Command) ReqAdmin() Command {
	c.AdminOnly = true
	return c
}
