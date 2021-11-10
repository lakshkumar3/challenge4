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
	settings.SetConfigName(flagEnvironment)
	settings.AddConfigPath(flagConfigPath)
	if err := settings.ReadInConfig(); err != nil {
		log.Fatal("Could not parse configuration file  " + flagConfigPath + "/" + flagEnvironment + ":" + err.Error())
		return
	}
	address := fmt.Sprint(flagEnvironment + "host:" + settings.GetString("REDIS_PORT"))
	opts := &redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	}
	warehouse.AddConnection("redis", redis.ConnectRedis(opts))
	err := server.StartServer()
	if err != nil {
		log.Error("server returing error " + err.Error())
	}

}
