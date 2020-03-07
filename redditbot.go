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
	NewCommand("reddit", false, RedditBot).Add()
}

// RedditBot creates a bot and grabs posts from a given subreddit
func RedditBot(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	// i feel like there's a better way to do this instead of
	// each time the command gets used
	bot, err := reddit.NewBotFromAgentFile("bot.agent", 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	subreddit := "/r/" + msgList[1]

	harvest, err := bot.Listing(subreddit, "")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if len(harvest.Posts) == 0 {
		s.ChannelMessageSend(m.ChannelID, "subreddit had no posts, try again.")
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
