package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("nsfw", false, nsfw).SetHelp("get access to the NSFW channel").Add()
}

func nsfw(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	member, err := s.GuildMember(config.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if ok := IsInArray(config.NSFWRoleID, member.Roles); ok {
		err := s.GuildMemberRoleRemove(config.GuildID, m.Author.ID, config.NSFWRoleID)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	} else {
		err := s.GuildMemberRoleAdd(config.GuildID, m.Author.ID, config.NSFWRoleID)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}
}
