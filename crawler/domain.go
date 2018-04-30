package crawler

import (
	"gopkg.in/yaml.v2"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"errors"
	"fmt"
)

// Domain is a single code hosting service.
type Domain struct {
	Id          string `yaml:"id"`
	Description string `yaml:"description"`
	ClientApi   string `yaml:"client-api"`
	URL         string `yaml:"url"`
	RateLimit struct {
		ReqH int `yaml:"req/h"`
		ReqM int `yaml:"req/m"`
	} `yaml:"rate-limit"`
	BasicAuth string `yaml:"basic-auth"`
}

func ReadAndParseDomains(domainsFile string, redisClient *redis.Client) ([]Domain, error) {
	// Open and read domains file list.
	data, err := ioutil.ReadFile(domainsFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in reading %s file: %v", domainsFile, err))
	}
	// Parse domains file list.
	domains, err := parseDomainsFile(data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in parsing %s file: %v", domainsFile, err))
	}
	log.Info("Loaded and parsed domains.yml")

	// Update the start URL if a failed one found in Redis.
	for _, domain := range domains {
		domain.updateStartURL(redisClient)
	}

	return domains, nil
}

// parseDomainsFile parses the domains file to build a slice of Domain.
func parseDomainsFile(data []byte) ([]Domain, error) {
	domains := []Domain{}

	// Unmarshal the yml in domains list.
	err := yaml.Unmarshal(data, &domains)
	if err != nil {
		return nil, err
	}

	return domains, nil
}

// updateStartURL checks if a repository list previously failed to be retrieved.
func (domain Domain) updateStartURL(redisClient *redis.Client) error {
	// Check if there is an URL that wasn't correctly retrieved.
	// URL.value="false" => set domain.URL to that one
	keys, err := redisClient.HKeys(domain.Id).Result()
	if err != nil {
		return err
	}

	// N launch. Check if some repo list was interrupted.
	for _, key := range keys {
		if redisClient.HGet(domain.Id, key).Val() == "failed" {
			log.Debugf("Found one interrupted URL. Starts from here: %s", key)
			domain.URL = key
		}
	}

	return nil
}

func (domain Domain) processAndGetNextURL(url string, repositories chan Repository) (string, error) {
	crawler := GetClientApiCrawler(domain.ClientApi)
	return crawler(domain, url, repositories)
}
