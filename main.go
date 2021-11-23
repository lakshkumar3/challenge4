package main

import (
	"challenge1/db"
	"challenge1/server"
	"challenge1/service"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"github.com/cakemarketing/snowbank/stores/redis"
	"github.com/cakemarketing/snowbank/warehouse"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pborman/getopt"
	"sync"
	"time"
)

// build specific variables
var (
	flagConfigPath  = "config"
	flagEnvironment = "local"
	wg              sync.WaitGroup
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
	host := settings.GetString("SQL_SERVER_HOST")
	port := settings.GetString("SQL_SERVER_PORT")
	databaseName := settings.GetString("DB_NAME")
	sqlUserName := settings.GetString("SQL_USERNAME")
	sqlPassword := settings.GetString("SQL_PASSWORD")
	sqlcoon, err := db.Connect(host, port, databaseName, sqlUserName, sqlPassword)
	if err != nil {
		log.Error("sql err   " + err.Error())
	}
	err = sqlcoon.Ping()
	if err != nil {
		log.Error("sql err   " + err.Error())
	}

	log.Info("sucessfully connected")
	warehouse.AddConnection("aurora", sqlcoon)
	log.Info("   • connecting to Redis ")

	address := fmt.Sprint(settings.GetString("REDIS_HOST") + ":" + settings.GetString("REDIS_PORT"))
	opts := &redis.Options{
		Addr:     address,
		Password: settings.GetString("REDIS_PASSWORD"),
		DB:       settings.GetInt("DB_NUM"),
	}
	warehouse.AddConnection("redis", db.ConnectRedis(opts))
	log.Info("   • connecting to Redis    ")

	errorchan := make(chan error)
	done := make(chan bool)
	BACKUP_INTERVAL_SECONDS := settings.GetInt("BACKUP_INTERVAL_SECONDS")
	ticker := time.NewTicker(time.Duration(BACKUP_INTERVAL_SECONDS) * time.Second)
	go func() {

		for {
			wg.Wait()
			<-ticker.C
			sqlcoon.Reopen()
			backup(errorchan, done)
		}

	}()
	go func() {
		for {
			select {
			case <-done:
				sqlcoon.Close()
				log.Info("backup done")
				wg.Done()
				continue
			case err = <-errorchan:
				log.Error("error " + err.Error())
			}
		}
	}()
	err = server.StartServer()
	if err != nil {
		log.Error("server returing error " + err.Error())
	}

}

func backup(errorChan chan<- error, done chan<- bool) {
	userService := service.UserService{}
	err := userService.SaveCache()
	if err != nil {
		log.Error("error while backing up" + err.Error())
		errorChan <- err
	} else {
		log.Info("backup completed")
		wg.Add(1)
		done <- true
		return
	}
}
