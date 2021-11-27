package main

import (
	"errors"
	"fmt"
	"nemesisbot/config"
	"nemesisbot/database"
	"os"
	"time"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
)

func main() {
	config.ReadConfig()

	setupDatabase()
	setupGalio()

}
func setupDatabase() {
	databasefolders := "sqlite/"
	databasename := "database.db"
	os.Mkdir(databasefolders, os.ModePerm)

	databasepath := databasefolders + databasename
	database.Init(databasepath)
	if _, err := os.Stat(databasepath); errors.Is(err, os.ErrNotExist) {
		database.CreateDatabase()
	}

	database.AddEntry("foo", "bar")

}

func setupGalio() {
	client := golio.NewClient(config.RiotToken,
		golio.WithRegion(api.RegionEuropeWest))
	summoner, _ := client.Riot.LoL.Summoner.GetByName("DreiAugenFlappe")
	matches, _ := client.Riot.LoL.Match.List(summoner.PUUID, 0, 5)

	for _, matchName := range matches {

		match, _ := client.Riot.LoL.Match.Get(matchName)
		fmt.Print(time.Unix(match.Info.GameStartTimestamp/1000, 0))
		fmt.Print(": ")
		participants := match.Info.Participants

		findNemesis(participants)

		fmt.Print("\n")
	}
}

func findNemesis(participants []*lol.Participant) int {
	for id, participant := range participants {
		fmt.Print(id)
		fmt.Print(" ")
		fmt.Print(participant.ChampionName + ", ")
	}
	return -1

}

//config.ReadConfig()
//dg, err := discordgo.New("Bot " + config.DiscordToken)
//if err != nil {
//	fmt.Println("error creating Discord session,", err)
//	return
//}
//
////eventhandler.Init()
////dg.AddHandler(eventhandler.VoiceChannelCreate)
//
//dg.Identify.Intents = discordgo.IntentsAll
//
//err = dg.Open()
//if err != nil {
//	fmt.Println("error opening connection,", err)
//	return
//}
//
//fmt.Println("Bot is now running.  Press CTRL-C to exit.")
//sc := make(chan os.Signal, 1)
//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
//<-sc
//
//dg.Close()
