package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Public variables
	DiscordToken   string
	RiotToken      string
	DiscordChannel string

	// Private variables
	config *configStruct
)

type configStruct struct {
	DiscordToken   string `json:"DiscordToken"`
	RiotToken      string `json:"RiotToken"`
	DiscordChannel string `json:"DiscordChannel"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	DiscordToken = config.DiscordToken
	RiotToken = config.RiotToken
	DiscordChannel = config.DiscordChannel

	return nil
}
