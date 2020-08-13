package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	NewCommand("fursona", false, fursona).SetHelp("create your own fursona! \n\npowered by: https://thisfursonadoesnotexist.com/" +
		"\n\nall images are created by an AI and non-copyrightable").Add()
}

func fursona(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	var paddedSeed string
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(99999)
	randomNumberString := strconv.Itoa(randomNumber)

	if x := countDigits(randomNumber); x <= 4 {
		paddedSeed = padNum(x, randomNumberString)
		em := NewEmbed().
			SetTitle("here's your fursona!").
			SetImage("https://thisfursonadoesnotexist.com/v2/jpgs-2x/seed" + paddedSeed + ".jpg").MessageEmbed

		s.ChannelMessageSendEmbed(m.ChannelID, em)
		return
	}

	em := NewEmbed().
		SetTitle("here's your fursona!").
		SetImage("https://thisfursonadoesnotexist.com/v2/jpgs-2x/seed" + randomNumberString + ".jpg").MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}

func countDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return
}

func padNum(count int, seed string) (paddedSeed string) {
	switch {
	case count == 4:
		paddedSeed = "0" + seed
	case count == 3:
		paddedSeed = "00" + seed
	case count == 2:
		paddedSeed = "000" + seed
	case count == 1:
		paddedSeed = "0000" + seed
	}
	return
}
