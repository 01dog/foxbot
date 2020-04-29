package main

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/anthonynsimon/bild/transform"
	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/primitive/primitive"
)

func init() {
	NewCommand("pri", false, pri).SetHelp("morph a user's avatar \n\nargs: `@user` `shape` `alpha` `rep`\n\n`shapes`: 0=combo," +
		"1=triangle, 2=rect, 3=ellipse, 4=circle, 5=rotatedrect, 6=beziers, 7=rotatedellipse, 8=polygon" +
		"\n\n`alpha`: 0-256\n\n`rep`:0-5 (works best with beziers, reps are hard on server resources so please be gentle)" +
		"\n\nuse no `@user` to morph your own avatar and no shape/alpha/rep for random values").Add()
}

func pri(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	if len(m.Mentions) == 0 {
		args := msgList[1:]
		avatar, err := GetAvatarImage(s, m.Author.ID)
		if err != nil {
			fmt.Println("error getting user avatar:", err)
			return
		}

		img := morphImage(s, m, avatar, args)
		filename := m.Author.ID + ".jpg"

		err = SaveImage(img, ".", filename)
		if err != nil {
			fmt.Println("error:", err)
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
				Title: m.Author.Username + "'s morphed avatar",
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://" + filename,
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: "made with primitive by fogleman",
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

	if len(m.Mentions) != 0 {
		args := msgList[2:]

		avatar, err := GetAvatarImage(s, m.Mentions[0].ID)
		if err != nil {
			fmt.Println("error getting user avatar:", err)
			return
		}

		img := morphImage(s, m, avatar, args)
		filename := m.Author.ID + ".jpg"

		err = SaveImage(img, ".", filename)
		if err != nil {
			fmt.Println("error:", err)
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
				Title: m.Mentions[0].Username + "'s morphed avatar",
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://" + filename,
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: "made with primitive by fogleman",
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
}

func morphImage(s *discordgo.Session, m *discordgo.MessageCreate, img image.Image, args []string) image.Image {
	s.ChannelMessageSend(m.ChannelID, "creating image, please wait...")
	rand.Seed(time.Now().UnixNano())

	img = transform.Resize(img, 256, 256, transform.Linear)
	bg := primitive.MakeColor(primitive.AverageImageColor(img))
	model := primitive.NewModel(img, bg, 512, runtime.NumCPU())
	argsInt, err := StrArrayToInt(args)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "error converting arguments to type int, try again")
	}

	if len(args) < 3 {
		for i := 0; i < 100; i++ {
			model.Step(primitive.ShapeType(0), 0, 0)
		}

		s.ChannelMessageSend(m.ChannelID, "image done!")
		return model.Context.Image()
	}

	for i := 0; i < 100; i++ {
		model.Step(primitive.ShapeType(argsInt[0]), argsInt[1], argsInt[2])
	}

	s.ChannelMessageSend(m.ChannelID, "image done!")
	return model.Context.Image()
}
