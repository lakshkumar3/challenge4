package main

import (
	"challenge1/server"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
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
		fmt.Printf("Could not parse configuration file '%s/%s': %v", flagConfigPath, flagEnvironment, err)
		log.Fatal("Could not parse configuration file  " + flagConfigPath + "/" + flagEnvironment + ":" + err.Error())
		return
	}
	err := server.Server()
	if err != nil {
		log.Error("server returing error " + err.Error())
	}
}
