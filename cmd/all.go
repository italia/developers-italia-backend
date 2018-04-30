package cmd

import (
	"fmt"
	"io/ioutil"
	"github.com/italia/developers-italia-backend/crawler"
	"github.com/italia/developers-italia-backend/metrics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(allCmd)
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Crawl publiccode.yml from hostings.",
	Long: `Start the crawler on every host written on hosting.yml file.
Beware! May take days to complete.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Init Prometheus for metrics.
		processedCounter := metrics.PrometheusCounter("repository_processed", "Number of repository processed.")

		// Open and read hosting file list.
		hostingFile := "hosting.yml"
		data, err := ioutil.ReadFile(hostingFile)
		if err != nil {
			panic(fmt.Sprintf("error in reading %s file: %v", hostingFile, err))
		}
		// Parse hosting file list.
		hostings, err := crawler.ParseHostingFile(data)
		if err != nil {
			panic(fmt.Sprintf("error in parsing %s file: %v", hostingFile, err))
		}
		log.Info("Loaded and parsed hosting.yml")

		// Initiate a channel of repositories.
		repositories := make(chan crawler.Repository)

		// Process each hosting service.
		for _, hosting := range hostings {
			go crawler.ProcessHosting(hosting, repositories)
		}

		// Process the repositories in order to retrieve publiccode.yml.
		crawler.ProcessRepositories(repositories, processedCounter)
	},
}
