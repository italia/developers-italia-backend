package crawler

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// Crawler is the interface for every specific crawler instances.
type Crawler interface {
	GetRepositories(url string, repositories chan Repository) (string, error)
}

// Process delegates the work to single hosting crawlers.
func Process(hosting Hosting, repositories chan Repository) {
	if hosting.ServiceInstance == nil {
		log.Warnf("Hosting %s is not available.", hosting.ServiceName)
		return
	}

	// Redis connection.
	redisClient, err := redisClientFactory("redis:6379")
	if err != nil {
		log.Error(err)
	}

	// Base starting URL.
	url := hosting.URL

	for {
		// Set the value of nextURL on redis to "failed".
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		err = redisClient.HSet(hosting.ServiceName, url, "failed").Err()
		if err != nil {
			log.Error(err)
		}
		log.Infof("%s saved on redis.", url)

		nextURL, err := hosting.ServiceInstance.GetRepositories(url, repositories)
		if err != nil {
			log.Errorf("error reading %s repository list: %v", url, err)
			close(repositories)
			return
		}
		time.Sleep(5 * time.Second)
		// If reached, the repository list was successfully retrieved.
		// Delete the repository url from redis.
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
		err = redisClient.HDel(hosting.ServiceName, url).Err()
		if err != nil {
			log.Error(err)
		}
		log.Infof("%s removed from redis at %s.", url, timestamp)
		time.Sleep(5 * time.Second)
		// Update url to nextURL.
		url = nextURL
	}

}
