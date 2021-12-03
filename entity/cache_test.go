package entity

import (
	"database/sql"
	"fmt"
	redis2 "github.com/go-redis/redis"
	"testing"
)

var (
	Scanfunc func(cursor uint64, match string, count int64) *redis2.ScanCmd
	SPopfunc func(key string) *redis2.StringCmd
	Execfunc func(query string, args ...interface{}) (sql.Result, error)
	SAddfunc func(key string, members ...interface{}) *redis2.IntCmd
)

type MockSql struct {
}

func (s MockSql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return Execfunc(query, args...)
}

type MockRedis struct {
}

func (c MockRedis) SPop(key string) *redis2.StringCmd {
	return SPopfunc(key)
}
func (c MockRedis) Scan(cursor uint64, match string, count int64) *redis2.ScanCmd {
	return Scanfunc(cursor, match, count)
}
func (c MockRedis) SAdd(key string, members ...interface{}) *redis2.IntCmd {
	return SAddfunc(key, members...)
}

type CacheMock struct{}

func TestCache_SaveEquationCache(t *testing.T) {
	redisdb := &MockRedis{}
	sqldb := &MockSql{}
	errNIl := true
	SPopfunc = func(key string) *redis2.StringCmd {
		var err error
		if errNIl {
			err = nil
			errNIl = false
		} else {
			err = redis2.Nil
		}

		var val string
		val = "{\"expresion\":\"2+3\",\"result\":\"5\",\"timestamp\":\"2021-12-02T20:32:00.1158939+05:00\"}"
		return redis2.NewStringResult(val, err)
	}
	Scanfunc = func(cursor uint64, match string, count int64) *redis2.ScanCmd {
		var err error = nil
		var val []string
		val = append(val, "laksh")
		return redis2.NewScanCmdResult(val, 0, err)
	}
	Execfunc = func(query string, args ...interface{}) (sql.Result, error) {
		return nil, nil
	}
	err := cache.SaveEquationCache(redisdb, sqldb)
	if err != nil {
		t.Fail()
	}

}
func TestNegativeCache_SaveEquationCache(t *testing.T) {
	redisdb := &MockRedis{}
	sqldb := &MockSql{}
	Scanfunc = func(cursor uint64, match string, count int64) *redis2.ScanCmd {
		err := fmt.Errorf("error 404")
		var val []string
		return redis2.NewScanCmdResult(val, 0, err)
	}
	err := cache.SaveEquationCache(redisdb, sqldb)
	if err == nil {
		t.Fail()
	}
	SPopfunc = func(key string) *redis2.StringCmd {
		err := redis2.Nil
		var val string
		return redis2.NewStringResult(val, err)
	}
	Scanfunc = func(cursor uint64, match string, count int64) *redis2.ScanCmd {
		var err error = nil
		var val []string
		val = append(val, "laksh")
		return redis2.NewScanCmdResult(val, 0, err)
	}
	err = cache.SaveEquationCache(redisdb, sqldb)
	if err != nil {
		t.Fail()
	}

	errNIl := true
	SPopfunc = func(key string) *redis2.StringCmd {
		//  err:=redis2.Nil

		var err error
		if errNIl {
			err = nil
			errNIl = false
		} else {
			err = redis2.Nil
		}

		var val string
		val = "{\"expresion\":\"2+3\",\"result\":\"5\",\"timestamp\":\"2021-12-02T20:32:00.1158939+05:00\"}"
		return redis2.NewStringResult(val, err)
	}
	Scanfunc = func(cursor uint64, match string, count int64) *redis2.ScanCmd {
		var err error = nil
		var val []string
		val = append(val, "laksh")
		return redis2.NewScanCmdResult(val, 0, err)
	}

	Execfunc = func(query string, args ...interface{}) (sql.Result, error) {
		return nil, fmt.Errorf("insert mock error")
	}
	SAddfunc = func(key string, members ...interface{}) *redis2.IntCmd {
		return redis2.NewIntResult(1, fmt.Errorf("mock sadd erorr"))
	}
	err = cache.SaveEquationCache(redisdb, sqldb)
	if err == nil {
		t.Fail()
	}

}
