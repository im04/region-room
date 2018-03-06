package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"time"
	"region-room/config"
)
var Db *xorm.Engine
var RedisPool *redis.Pool

func initDb() {
	var err error
	Db, err = xorm.NewEngine(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		panic("error when connect database!,err:" + err.Error())
	}
	Db.SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	Db.SetMaxOpenConns(config.DBConfig.MaxOpenConns)
}

func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	initDb()
	initRedis()
}