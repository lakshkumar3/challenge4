package entity

import (
	"challenge1/db"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/warehouse"
)

var (
	cache Cacher
)

func init() {
	cache = &Cache{}
}

type Cache struct {
}
type Cacher interface {
	SaveEquationCache() error
}

func (*Cache) SaveEquationCache() error {
	opts := &warehouse.Options{}
	msSqlinterface := warehouse.GetConnection("aurora", opts)
	sqldb, ok := msSqlinterface.(*db.Mysql)
	if !ok {
		log.Error("error while parsing")
	}
	err := sqldb.Ping()
	if err != nil {
		log.Error(err.Error())
	}
	redisinterface := warehouse.GetConnection("redis", opts)
	redisdb, ok := redisinterface.(*db.RedisCache)
	if !ok {
		log.Fatal("redis parse error")
	}
	redisdb.MU.Lock()
	users, _, err := redisdb.Scan(0, "*", 100).Result()
	if err != nil {
		log.Error("error while scaning redis " + err.Error())
		return err
	}

	for _, user := range users {
		allEqByUser, err := redisdb.SMembers(user).Result()

		if err != nil {
			log.Error("user :" + user + "   err : " + err.Error())
			//		return err
		} else {
			var equation Equation
			for _, equationJson := range allEqByUser {
				err = json.Unmarshal([]byte(equationJson), &equation)
				if err != nil {
					log.Error("while marshling " + err.Error())
				}
				err = insertEquationToSQL(equation, sqldb, user)
				if err != nil {
					log.Error("error while saving data to sql  details:" + err.Error())
				}
			}
			log.Info(fmt.Sprintf("user :%s %v", user, allEqByUser))
		}
	}
	redisdb.FlushAll()
	redisdb.MU.Unlock()
	return nil
}
func insertEquationToSQL(equation Equation, db *db.Mysql, user string) error {
	insertQuery := "insert into Equations (UserName,Expression,Result,Eqtimestamp) VALUES (?, ?, ?, ?);"
	result, err := db.Exec(insertQuery, user, equation.Expresion, equation.Result, equation.Timestamp)

	if err != nil {
		return err
		log.Error(err.Error())
	}
	log.Info(fmt.Sprintf("result  %v", result))
	return nil
}
