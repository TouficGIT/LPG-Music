package framework

import (
	"encoding/json"
	"io/ioutil"

	"Work.go/LPG-Bot/LPGMusic/print"
)

type Config struct {
	Token       string `json:"Token"`
	BotPrefix   string `json:"BotPrefix"`
	YtbToken    string `json:"YtbToken"`
	PurgeTime   int    `json:"Purgetime"`
	UseSharding bool   `json:"use_sharding"`
	ShardId     int    `json:"shard_id"`
	ShardCount  int    `json:"shard_count"`
}

func LoadConfig(filename string) *Config {
	print.InfoLog("[INFO] Reading from config file ...", "[SERVER]")
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		print.CheckError("[ERROR] Reading config.json file", "[Server]", err)
		return nil
	}
	var conf Config
	err = json.Unmarshal(body, &conf)
	if err != nil {
		print.CheckError("[ERROR] Unmarshal config.json data", "[Server]", err)
		return nil
	}
	return &conf
}
