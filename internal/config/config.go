package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"ses_back/internal/models"
)

const YamlPath = "/home/starman/GolandProjects/ses_back/config.yaml"

var Config models.AppConfig

func ReadConfig() {
	file, err := os.ReadFile(YamlPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		panic(err)
	}
}
