package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env      string `yaml:"env"`
	Addr     string `yaml:"address"`
	DataBase struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"dbname"`
	}
}

func ReadConfig(file string) (*Config, error) {
	var c Config
	buff, err := os.ReadFile(file)
	if err != nil {
		return &c, err
	}
	err = yaml.Unmarshal(buff, &c)
	return &c, err
}
