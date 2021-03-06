package main

import (
	"encoding/json"
	"log"

	"github.com/louisevanderlith/mango/util"
)

type DomainSetting struct {
	Address string
	Name    string
	Type    string
}

type Settings []DomainSetting

func loadSettings() *Settings {
	dbConfPath := util.FindFilePath("domains.json", "conf")
	content := util.GetFileContent(dbConfPath)

	var settings *Settings
	err := json.Unmarshal(content, &settings)

	if err != nil {
		log.Print("loadSettings: ", err)
	}

	return settings
}
