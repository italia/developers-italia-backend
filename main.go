package main

import (
	"fmt"

	"github.com/italia/developers-italia-backend/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/italia/developers-italia-backend/crawler"
)

func main() {
	log.SetLevel(log.DebugLevel)

	// Read configurations.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reding config file: %s", err))
	}

	// Register client APIs.
	crawler.RegisterClientApis()

	cmd.Execute()
}
