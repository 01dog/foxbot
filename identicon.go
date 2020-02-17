package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("identicon", false, identicon).SetHelp("generate your identicon" +
		"\n\nidenticons are visual representations of hashes. each identicon is unique to you!" +
		"\nmore info here: https://barro.github.io/2018/02/avatars-identicons-and-hash-visualization/").Add()
}

func identicon(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	img := GenerateIdenticon(m.Author.ID)
	filename, err := RenderIdenticon(img, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "there was an error creating your identicon, try again!")
		return
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer f.Close()

	// manually creating this embed to attach the image
	// could probably add file attach to embed.go later
	em := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: m.Author.Username + "'s identicon",
			Image: &discordgo.MessageEmbedImage{
				URL: "attachment://" + filename,
			},
		},
		Files: []*discordgo.File{
			&discordgo.File{
				Name:   filename,
				Reader: f,
			},
		},
	}
	s.ChannelMessageSendComplex(m.ChannelID, em)
}
