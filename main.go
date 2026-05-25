package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
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
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("user-data-dir", cp),
	)
	fmt.Println(len(opts))
	fmt.Println("Profile Directory : "+cp, stupidfuckingreasonlessvariable)
	fuckingAllocShit, dontbreakplease := chromedp.NewExecAllocator(context.Background(), opts...)
	defer dontbreakplease()
	ctext, dontbreakpleaseb := chromedp.NewContext(fuckingAllocShit)
	defer dontbreakpleaseb()
	browsererr := chromedp.Run(ctext,
		chromedp.Navigate("https://instagram.com/reels/"),
	)
	if browsererr != nil {
		log.Fatal(browsererr)
	}
	fmt.Println("Make sure instagram is logged in and then press enter to continue")
	fmt.Scanln()
	var URLbk string
	stopLoopChan := make(chan bool)
	go func() {
		fmt.Println("Starting loop, press enter to stop from here")
		fmt.Scanln()
		stopLoopChan <- true
	}()
	for {
		var URL string

		vidfinderr := chromedp.Run(ctext,
			chromedp.KeyEvent(kb.ArrowDown),
			chromedp.Sleep(5*time.Second),
			chromedp.Location(&URL),
		)
		if vidfinderr != nil {
			log.Fatal(vidfinderr)
		}
		if URL == URLbk {
			fmt.Println("An error has occured resulting in a failure to scroll, refreshing the page now.")
			vidfinderr := chromedp.Run(ctext,
				chromedp.Navigate("https://instagram.com/reels/"),
				chromedp.Sleep(10*time.Second),
				chromedp.Location(&URL),
			)
			if vidfinderr != nil {
				log.Fatal(vidfinderr)
			}
		}
		posterr, fuck := dg.ChannelMessageSend(channel, strings.Replace(URL, "instagram", "vxinstagram", 1))
		if posterr != nil {
			log.Println("Failed to send message:", posterr)
			log.Println(fuck)
		}
		time.Sleep(12 * time.Second)
		time.Sleep(time.Duration(rand.IntN(5)) * time.Second)

		URLbk = URL
		select {
		case <-stopLoopChan:
			return
		default:
		}
	}

}
