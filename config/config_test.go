package config_test

import (
	"errors"
	"testing"

	"github.com/raymondwongso/kapyrpg/config"
	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	tests := map[string]struct {
		envPath        string
		expectedConfig config.Config
		expectedErr    error
	}{
		"valid env": {
			envPath: "./test/valid.env",
			expectedConfig: config.Config{
				BotToken:      "test_bot_token",
				CommandPrefix: "!",
				ConfuseEmoji:  "😕",
			},
		},
		"env with empty bot token": {
			envPath:        "./test/empty_bot_token.env",
			expectedConfig: config.Config{},
			expectedErr:    errors.New("discord bot token is empty"),
		},
		"env with empty command prefix": {
			envPath: "./test/empty_command_prefix.env",
			expectedConfig: config.Config{
				BotToken:      "test_bot_token",
				CommandPrefix: "!",
				ConfuseEmoji:  "😕",
			},
		},
		"env with empty confuse emoji": {
			envPath: "./test/empty_confuse_emoji.env",
			expectedConfig: config.Config{
				BotToken:      "test_bot_token",
				CommandPrefix: "!",
				ConfuseEmoji:  "😕",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg, err := config.Load(tc.envPath)
			assert.Equal(t, tc.expectedConfig, cfg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
