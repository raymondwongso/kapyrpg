package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/raymondwongso/kapyrpg/config"
	"github.com/raymondwongso/kapyrpg/handler"
	"github.com/rs/zerolog/log"
)

func bot() error {
	log.Info().Msg("loading configuration files...")
	config, err := config.Load()
	if err != nil {
		log.Fatal().
			Err(err).Msg("error loading configuration files")
	}

	log.Info().
		Str("config", fmt.Sprintf("%+v\n", config)).Msg("configuration file succesfully loaded.")

	log.Info().Msg("initializing discord session...")
	session, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatal().Err(err).Msg("error initializing discord session")
	}

	// helpService := help_service.New()
	handler := handler.New(session, config)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msg("preparing bot...")
		if err := handler.Run(); err != nil {
			log.Fatal().Err(err).Msg("error running bot handler")
		}
	}()

	<-done

	log.Info().Msg("stopping bot...")
	log.Info().Msg("bot stopped, see you next time")
	return nil
}
