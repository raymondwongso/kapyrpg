package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	defaultCommandPrefix = "!"
	defaultConfuseEmoji  = "😕"
)

type Config struct {
	BotToken      string `env:"BOT_TOKEN"`
	CommandPrefix string `env:"COMMAND_PREFIX"`
	ConfuseEmoji  string `env:"CONFUSE_EMOJI"`
}

// Load loads configuration file and return Config object
func Load(filenames ...string) (Config, error) {
	err := godotenv.Overload(filenames...)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		BotToken:      os.Getenv("BOT_TOKEN"),
		CommandPrefix: os.Getenv("COMMAND_PREFIX"),
		ConfuseEmoji:  os.Getenv("CONFUSE_EMOJI"),
	}

	if cfg.BotToken == "" {
		return Config{}, errors.New("discord bot token is empty")
	}
	if cfg.CommandPrefix == "" {
		cfg.CommandPrefix = defaultCommandPrefix
	}
	if cfg.ConfuseEmoji == "" {
		cfg.ConfuseEmoji = defaultConfuseEmoji
	}

	return cfg, nil
}
