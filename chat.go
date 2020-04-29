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
	NewCommand("nyoom", false, nyoom).Add()
	NewCommand("dmc3", false, dmc3).Add()
}

func dogMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"björk", "*woof*", "*grrrrr*", "*wags tail*", "^w^",
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

func nyoom(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	em := NewEmbed().
		SetFooter("nyoooooooooooooooooooooooooooooooooooooooooooooom").
		SetImage("https://media.discordapp.net/attachments/562589611310841858/676315314547458062/nyooooooom.gif").MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}

func dmc3(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"Errr, I think you will laughing with your cola upsiding down watching this.", "I don't think anyone of you may notice this, tricky thing?",
		"I bet many of your mind goes, boooooooom...", "Yeaa yeaa, I know you're laughing off me, bear me, I know I missed, hello~ low spec PC." +
			" What can you expect? Gimme your razor of silence!", "What else haha///",
	}
	rand.Seed(time.Now().UnixNano())
	messageChosen := messages[rand.Intn(len(messages))]

	s.ChannelMessageSend(m.ChannelID, messageChosen)
}
