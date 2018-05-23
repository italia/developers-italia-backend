package cmd

import (
	"strconv"
	"sync"
	"time"

	"github.com/italia/developers-italia-backend/crawler"
	"github.com/italia/developers-italia-backend/metrics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var iPAToCrawl string

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.Flags().StringVarP(&iPAToCrawl, "ipa", "i", "", "Crawl a single ipa from whitelist.yml.")
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl publiccode.yml from domains in whitelist.",
	Long:  `Start the crawler on whitelist.yml file.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Elastic connection.
		elasticClient, err := crawler.ElasticClientFactory(
			viper.GetString("ELASTIC_URL"),
			viper.GetString("ELASTIC_USER"),
			viper.GetString("ELASTIC_PWD"))
		if err != nil {
			panic(err)
		}

		// Read and parse list of domains.
		domainsFile := "domains.yml"
		domains, err := crawler.ReadAndParseDomains(domainsFile)
		if err != nil {
			panic(err)
		}

		// Read and parse the whitelist.
		whitelistFile := "whitelist.yml"
		whitelist, err := crawler.ReadAndParseWhitelist(whitelistFile)
		if err != nil {
			panic(err)
		}

		// Initiate a channel of repositories.
		repositories := make(chan crawler.Repository, 1000)
		// Prepare WaitGroup.
		var wg sync.WaitGroup

		// Index for actual process.
		index := strconv.FormatInt(time.Now().Unix(), 10)

		// Register Prometheus metrics.
		metrics.RegisterPrometheusCounter("repository_processed", "Number of repository processed.", index)
		metrics.RegisterPrometheusCounter("repository_file_saved", "Number of file saved.", index)
		metrics.RegisterPrometheusCounter("repository_file_indexed", "Number of file indexed.", index)
		// Uncomment when validating publiccode.yml
		//metrics.RegisterPrometheusCounter("repository_file_saved_valid", "Number of valid file saved.", index)

		// Process every item in whitelist.
		for _, pa := range whitelist {
			wg.Add(1)
			// If iPAToCrawl is empty crawl all domains, otherwise crawl only the one with CodiceIPA equals to iPAToCrawl.
			if (iPAToCrawl == "") || (iPAToCrawl != "" && pa.CodiceIPA == iPAToCrawl) {
				go crawler.ProcessPA(pa, domains, repositories, index, &wg)
			}
		}

		// Start the metrics server.
		go metrics.StartPrometheusMetricsServer()

		// WaitingLoop check and close the repositories channel
		go crawler.WaitingLoop(repositories, index, &wg, elasticClient)

		// Process the repositories in order to retrieve the file.
		crawler.ProcessRepositories(repositories, index, &wg, elasticClient)

	}}