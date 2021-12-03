package entity

import (
	"challenge1/db"
	"encoding/json"
	"github.com/apex/log"
	redis2 "github.com/go-redis/redis"
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
	SaveEquationCache(redisdb db.RedisCacher, sqldb db.EquationSaver) error
}

func (*Cache) SaveEquationCache(redisdb db.RedisCacher, sqldb db.EquationSaver) error {

	users, _, err := redisdb.Scan(0, "*", 100).Result()
	if err != nil {
		log.Error("error while scaning redis " + err.Error())
		return err
	}

	for _, user := range users {

		if err != nil {
			log.Error("user :" + user + "   err : " + err.Error())
			//		return err
		} else {
			var equation Equation
			for {
				equationJson, err := redisdb.SPop(user).Result()
				if err != nil {
					if err == redis2.Nil {
						break
					}
					log.Error("error occured while Pop " + err.Error())
				}

				err = json.Unmarshal([]byte(equationJson), &equation)

				if err != nil {
					log.Error("while marshling " + err.Error())
					return err
				}
				err = insertEquationToSQL(equation, sqldb, user)
				if err != nil {
					_, err := redisdb.SAdd(user, equationJson).Result()
					ctx := log.WithFields(log.Fields{
						"error":      err.Error(),
						"key user":   user,
						"value data": equationJson,
					})
					if err != nil {
						ctx.Error("not able to add data back to redis data lost   ")
					}
					log.Error("error while saving data to sql  details:" + err.Error())
					return err
				}
			}
		}
	}
	return nil
}
func insertEquationToSQL(equation Equation, db db.EquationSaver, user string) error {
	insertQuery := "insert into Equations (UserName,Expression,Result,Eqtimestamp) VALUES (?, ?, ?, ?);"
	_, err := db.Exec(insertQuery, "laksh_test", equation.Expresion, equation.Result, equation.Timestamp)

	if err != nil {
		log.Error(err.Error())
		return err
	}
	//log.Info(fmt.Sprintf("result  %v", result))
	return nil
}
