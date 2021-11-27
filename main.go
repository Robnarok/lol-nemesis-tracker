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
}

func setupGalio() {
	client := golio.NewClient(config.RiotToken,
		golio.WithRegion(api.RegionEuropeWest))
	checkMatchhistory(client, "DreiAugenFlappe", "Teemo")
}

func checkMatchhistory(client *golio.Client, summonerToCheck string, nemesisName string) {

	matchesToCheck := fetchMatchList(client, summonerToCheck)
	fmt.Print(matchesToCheck)

	for _, matchName := range matchesToCheck {
		match, _ := client.Riot.LoL.Match.Get(matchName)
		fmt.Print(time.Unix(match.Info.GameStartTimestamp/1000, 0))
		fmt.Print(": ")
		participants := match.Info.Participants
		nemesisID := findNemesis(participants, nemesisName)
		if nemesisID >= 0 {
			fmt.Printf(participants[nemesisID].SummonerName)
		} else {
			fmt.Printf("Kein Nemesis im Match")
		}
		fmt.Print("\n")
		database.AddEntry(fmt.Sprintf("%d", match.Info.GameID), fmt.Sprint(match.Info.GameStartTimestamp))
	}
}

func fetchMatchList(client *golio.Client, summonerToCheck string) []string {
	summoner, _ := client.Riot.LoL.Summoner.GetByName(summonerToCheck)
	checked_matches := database.GetAllEntrys()
	matches, _ := client.Riot.LoL.Match.List(summoner.PUUID, 0, 5)
	matchesToCheck := make([]string, 0)
	foobar := true

	for _, matchFromLastMatches := range matches {
		for _, matchFromAllMatches := range checked_matches {
			if "EUW1_"+matchFromAllMatches.Match == matchFromLastMatches {
				foobar = false
			}
		}
		if foobar {
			matchesToCheck = append(matchesToCheck, matchFromLastMatches)
			foobar = true
		}
	}
	return matchesToCheck
}

func findNemesis(participants []*lol.Participant, nemesisName string) int {
	for id, participant := range participants {
		if participant.ChampionName == nemesisName {
			return id
		}
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
