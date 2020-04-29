package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/turnage/graw/reddit"
)

// TODO: clean this up a bit

func init() {
	NewCommand("reddit", false, RedditBot).SetHelp("get a random post from a subreddit\n\n args: `subreddit`").Add()
}

// RedditBot creates a bot and grabs posts from a given subreddit
func RedditBot(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	// i feel like there's a better way to do this instead of
	// each time the command gets used
	bot, err := reddit.NewBotFromAgentFile("bot.agent", 0)
	if err != nil {
		msg := fmt.Sprintf("error creating reddit bot: %s", err)
		s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	if len(msgList) < 2 {
		s.ChannelMessageSend(m.ChannelID, "i need a subreddit to search.")
		return
	}

	msgList = msgList[0:2]
	subreddit := "/r/" + msgList[1]
	harvest, err := bot.Listing(subreddit, "")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if len(harvest.Posts) == 0 {
		msg := fmt.Sprintf("no results found for %s, try another subreddit", subreddit)
		s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomPost := harvest.Posts[rand.Intn(len(harvest.Posts))]

	if randomPost.NSFW {
		s.ChannelMessageSend(m.ChannelID, "post is NSFW, try again.")
		return
	}

	em := NewEmbed().
		SetTitle(randomPost.Title).
		SetImage(randomPost.URL).MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, em)
}
