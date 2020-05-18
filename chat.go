package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("dog", false, dogMsg).Add()
	NewCommand("cat", false, catMsg).Add()
	NewCommand("fox", false, foxMsg).Add()
	NewCommand("mode", false, modeMsg).Add()
	NewCommand("nyoom", false, nyoom).Add()
	// NewCommand("quote", false, quote).Add()
}

func dogMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"björk", "*woof*", "*grrrrr*", "*wags tail*", "^w^",
	}
	s.ChannelMessageSend(m.ChannelID, messages[GenRandomNum(len(messages))])
}

func catMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		":3", "*meow*", "*purr*", "*knocks shit off the counter*", ":rage:",
		"**wow**", "*murr*", "*mjau*", "oh fuck meow",
	}
	s.ChannelMessageSend(m.ChannelID, messages[GenRandomNum(len(messages))])
}

func foxMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"AAAAAAAAAAA", "HIYAA", "blip",
	}
	s.ChannelMessageSend(m.ChannelID, messages[GenRandomNum(len(messages))])
}

func modeMsg(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	messages := []string{
		"dont care", "mega gay", "mode", "*yiffs*",
		"nice", "honestly vore isnt that bad", "*murrs*",
		"I SUMMON MY WARDS",
		"uncertain", "try again later", "D OMEGALUL C", "losing mmr", "fuck",
	}
	s.ChannelMessageSend(m.ChannelID, "mode "+messages[GenRandomNum(len(messages))])
}

func nyoom(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	em := NewEmbed().
		SetFooter("nyoooooooooooooooooooooooooooooooooooooooooooooom").
		SetImage("https://media.discordapp.net/attachments/562589611310841858/676315314547458062/nyooooooom.gif").MessageEmbed
	s.ChannelMessageSendEmbed(m.ChannelID, em)
}
