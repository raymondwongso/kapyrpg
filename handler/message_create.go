package handler

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var (
	BadCommand = errors.New("command unknown")
	cmdRegex   = regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
)

func (h *handler) MessageCreateHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ignore message from bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}

	args, err := h.messageCreateArgs(message.Content)
	if err != nil {
		switch err {
		case BadCommand:
			// warn bad command then silently react with confuse emoji
			log.Warn().Str("message.content", message.Content).Msg("bad command found")
			_ = h.reactConfuse(session, message.ChannelID, message.ID)
			return
		default:
			// unknown error
			log.Error().Err(err).Msg("error parsing args")
		}
	}

	var msgSend *discordgo.MessageSend

	switch args[0] {
	case "help":
		// msgSend, err = h.helpService.ListCommand(ctx)
		err = errors.New("help service is not immplemented")
	default:
		// command has valid format, but unknown. silently react with confuse emoji
		_ = h.reactConfuse(session, message.ChannelID, message.ID)
		return
	}

	if err != nil {
		log.Error().Err(err).Strs("args", args).Msg("error handling message create handler")
		return
	}

	_, err = session.ChannelMessageSendComplex(message.ChannelID, msgSend)
	if err != nil {
		log.Error().Err(err).Msg("error sending message feedback for MessageCreateHandler")
		return
	}
}

func (h *handler) messageCreateArgs(cmd string) ([]string, error) {
	if !h.isPrefixValid(cmd) {
		return nil, BadCommand
	}

	cmd = h.cleanPrefix(cmd)
	args := cmdRegex.FindAllString(cmd, -1)

	return args, nil
}
