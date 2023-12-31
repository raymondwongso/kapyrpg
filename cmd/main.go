package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "kapyrpg",
		Short: "KapyRPG Discord Bot",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Run KapyRPG Discord bot",
		RunE: func(cmd *cobra.Command, args []string) error {
			return bot()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Fatal().
			Err(err).Msg("[MAIN] error executing command")
	}
}
