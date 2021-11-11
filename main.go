package main

import (
	"challenge1/server"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"github.com/cakemarketing/snowbank/stores/redis"
	"github.com/cakemarketing/snowbank/warehouse"
	"github.com/pborman/getopt"
)

// build specific variables
var (
	flagConfigPath  = "config"
	flagEnvironment = "local"
)

func init() {
	getopt.StringVarLong(&flagConfigPath, "config-directory", 'p', "path to the config file")
	getopt.StringVarLong(&flagEnvironment, "environment", 'e', "environment of running instance")
	getopt.Parse()
	if len(getopt.Args()) > 0 {
		flagEnvironment = getopt.Arg(0)
	}
}

func main() {
	// parse the environment file
	ctx := log.WithFields(log.Fields{
		"Path":     flagConfigPath,
		"FileName": flagEnvironment,
	})
	settings.SetConfigName(flagEnvironment)
	settings.AddConfigPath(flagConfigPath)
	if err := settings.ReadInConfig(); err != nil {
		ctx.WithError(err).Error("Could not parse configuration file  ")
		return
	}
	address := fmt.Sprint(settings.GetString("REDIS_HOST") + ":" + settings.GetString("REDIS_PORT"))
	opts := &redis.Options{
		Addr:     address,
		Password: settings.GetString("REDIS_PASSWORD"),
		DB:       settings.GetInt("DB_NUM"),
	}
	warehouse.AddConnection("redis", redis.ConnectRedis(opts))
	log.Info("   • connecting to Redis    ")
	err := server.StartServer()
	if err != nil {
		log.Error("server returing error " + err.Error())
	}

}
