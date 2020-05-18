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

// func quote(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
// 	// TODO: move this to a file or something so it doesn't get so cluttered and gross
// 	messages := []string{
// 		"Errr, I think you will laughing with your cola upsiding down watching this.", "I don't think anyone of you may notice this, tricky thing?",
// 		"I bet many of your mind goes, boooooooom...", "Yeaa yeaa, I know you're laughing off me, bear me, I know I missed, hello~ low spec PC." +
// 			" What can you expect? Gimme your razor of silence!", "What else haha///",
// 		"I am one of three Elites who thumbed this down. My criterion? Glad you asked: Firstly, I was perturbed by the lack of \"verve.\" The" +
// 			" \"feel\" is there, but where is the \"presence?\" The subtle \"flow\" of the opus is interrupted by that jarring \"thrust\"" +
// 			" that Arvo inscrutably feels necessary to dribble in, not unlike an overly liberal sprinkling of artisanal pepper on your Kobe" +
// 			" \"filet mignon\" by a gauche \"maître d'hôtel,\" deaf to your cries of \"When! When!\"\n\n" +
// 			" © 2012 brighton dechienne All Rights Reserved",
// 		"dog just put his dick all over my face man\nthis is why I'm gay", "I'm afraid of holding babys because every time I do I can't help but" +
// 			" Imagine how cute and sexy they would turn in the future.", "my dog is licking his fat cock like I won't do it for him :rolling_eyes::rolling_eyes:",
// 	}
//
//

// 	s.ChannelMessageSend(m.ChannelID, messages[GenRandomNum(len(messages))])
// }
