package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("dog", false, dogMsg).Add()
	NewCommand("cat", false, catMsg).Add()
	NewCommand("fox", false, foxMsg).Add()
	NewCommand("mode", false, modeMsg).Add()
}

func dogMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"björk", "*woof*", "*grrrrr*", "*wags tail*", "*shits on the floor*",
		"^w^",
	}
	rand.Seed(time.Now().UnixNano())
	messageChosen := messages[rand.Intn(len(messages))]

	s.ChannelMessageSend(m.ChannelID, messageChosen)
}

func catMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		":3", "*meow*", "*purr*", "*knocks shit off the counter*", ":rage:",
		"**wow**", "*murr*", "*mjau*", "oh fuck meow",
	}
	rand.Seed(time.Now().UnixNano())
	messageChosen := messages[rand.Intn(len(messages))]

	s.ChannelMessageSend(m.ChannelID, messageChosen)
}

func foxMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"AAAAAAAAAAA", "HIYAA", "blip",
	}
	rand.Seed(time.Now().UnixNano())
	messageChosen := messages[rand.Intn(len(messages))]

	s.ChannelMessageSend(m.ChannelID, messageChosen)
}

func modeMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"dont care", "mega gay", "mode", "*yiffs*",
		"nice", "honestly vore isnt that bad", "*murrs*",
		"I SUMMON MY WARDS",
		"uncertain", "try again later", "D OMEGALUL C", "losing mmr", "fuck",
	}
	rand.Seed(time.Now().UnixNano())
	messageChosen := messages[rand.Intn(len(messages))]

	s.ChannelMessageSend(m.ChannelID, "mode "+messageChosen)
}
