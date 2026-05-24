package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		fmt.Println("Failed to load env file, make sure one is in the current directory.")
	}
	token := os.Getenv("BOT_TOKEN")
	channel := os.Getenv("BOT_CHANNEL")
	dg, sesserr := discordgo.New("Bot " + token)
	if sesserr != nil {
		fmt.Println("Session failure")
	}
	openerr := dg.Open()
	if openerr != nil {
		fmt.Println("Failed to open connection" + openerr.Error())
	}
	dg.ChannelMessageSend(channel, "DeRail is active and searching.")
	time.Sleep(2 * time.Second)
	cp, stupidfuckingreasonlessvariable := filepath.Abs("./cp")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-blink-features=AutomationControlled", true),
		chromedp.Flag("headless", false), // Headful is often less suspicious to basic anti-bot systems
	)
	fmt.Println(len(opts))
	fmt.Println("Profile Directory : "+cp, stupidfuckingreasonlessvariable)

}
