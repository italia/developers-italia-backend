package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

const domainsFile = "domains.yml"

var rootCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawler is a crawler for publiccode.yml file.",
	Long: `A Fast and Robust publiccode.yml file crawler.
Complete documentation is available at https://github.com/italia/developers-italia-backend`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute is the entrypoint for cmd package Cobra.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
