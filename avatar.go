package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("avatar", false, userAvatar).SetHelp("args: `@user`\nreturns the user's avatar").Add()
}

func userAvatar(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	if len(msgList) == 1 {
		getAvatar(s, m, m.Author.ID)
		return
	}
	if len(m.Mentions) != 0 {
		getAvatar(s, m, m.Mentions[0].ID)
		return
	}

	s.ChannelMessageSend(m.ChannelID, "user not found")
}

func getAvatar(s *discordgo.Session, m *discordgo.MessageCreate, userID string) {
	user, err := UserDetails(s, userID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "error finding the user, try again!")
		return
	}

	em := NewEmbed().
		SetTitle(user.Username + "'s avatar").
		SetImage(user.AvatarURL("2048")).MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}
