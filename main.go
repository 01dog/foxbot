package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var config = Configure()

func main() {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer dg.Close()
	dg.AddHandler(EvalMessage)
	err = dg.Open()
	if err != nil {
		fmt.Println("no connection:", err)
		return
	}

	dg.UpdateStatus(0, config.Prefix+"help")
	fmt.Println("bot running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("exiting")
}
