package models

import (
	"fmt"
	"os"
)

type AppConfigModel struct {
	ServerConfigModel struct {
		Host string
		Port int32
		Name string
	} `yaml:"server"`

	DatabaseConfigModel struct {
		Url      string
		Host     string
		Username string
		Password string
	} `yaml:"database"`
}

var (
	serverConfig   AppConfigModel
	databaseConfig *AppConfigModel
)

func GetServerConfig() *AppConfigModel {
	serverConfig := AppConfigModel{}
	fp, err := os.Open("config.yaml")
	if err == nil {
		// Means no err
		fmt.Println("Trting to read file:", fp)

	}

	return &serverConfig
}
