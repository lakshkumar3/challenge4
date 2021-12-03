package entity

import (
	"fmt"
	redis2 "github.com/go-redis/redis"
	"testing"
	"time"
)

func TestEquation_SaveEquation(t *testing.T) {
	redisdb := &MockRedis{}
	equation := Equation{
		Expresion: "1+2",
		Result:    "3",
		Timestamp: time.Now(),
	}
	SAddfunc = func(key string, members ...interface{}) *redis2.IntCmd {
		return redis2.NewIntResult(1, nil)
	}

	user := User{}
	user.SetName("laksh")
	err := equation.SaveEquation(user, redisdb)
	if err != nil {
		t.Fail()
	}

}
func TestNegativeEquation_SaveEquation(t *testing.T) {

	redisdb := &MockRedis{}
	equation := Equation{}
	SAddfunc = func(key string, members ...interface{}) *redis2.IntCmd {
		return redis2.NewIntResult(1, fmt.Errorf("mock sadd erorr"))
	}

	user := User{}
	user.SetName("laksh")
	err := equation.SaveEquation(user, redisdb)
	if err == nil {
		t.Fail()
	}

}
