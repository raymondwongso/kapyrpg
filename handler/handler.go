package handler

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/raymondwongso/kapyrpg/config"
	"github.com/rs/zerolog/log"
)

type handler struct {
	session *discordgo.Session
	config  config.Config
}

func New(
	session *discordgo.Session,
	cfg config.Config,
) *handler {
	return &handler{
		session: session,
		config:  cfg,
	}
}

func (h *handler) Run() error {
	h.session.AddHandler(h.MessageCreateHandler)

	err := h.session.Open()
	if err != nil {
		log.Error().Err(err).Msg("error opening discord session")
		return err
	}
	defer h.session.Close()

	log.Info().Msg("bot is up and running")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	return nil
}

func (h *handler) isPrefixValid(cmd string) bool {
	return len(cmd) >= len(h.config.CommandPrefix) || cmd[:len(h.config.CommandPrefix)] != h.config.CommandPrefix
}

func (h *handler) cleanPrefix(cmd string) string {
	return strings.TrimSpace(cmd[len(h.config.CommandPrefix):])
}

func (h *handler) reactConfuse(session *discordgo.Session, channelID, messageID string) error {
	err := session.MessageReactionAdd(channelID, messageID, h.config.ConfuseEmoji)
	if err != nil {
		log.Warn().Err(err).Msg("error adding confuse reaction")
	}

	return nil
}
