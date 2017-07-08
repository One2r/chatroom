package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

//RedisConnPool redis 连接池
var RedisConnPool *redis.Pool

func init() {
	maxIdle, _ := beego.AppConfig.Int("redis_pool_max_idle")
	idleTimeout, _ := beego.AppConfig.Int("redis_pool_idle_timeout")
	addr := beego.AppConfig.String("redis_host") + ":" + beego.AppConfig.String("redis_port")
	RedisConnPool = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				beego.Info(err)
				return nil, err
			}
			if _, err := c.Do("AUTH", beego.AppConfig.String("redis_password")); err != nil {
				c.Close()
				beego.Info(err)
				return nil, err
			}
			db, _ := beego.AppConfig.Int("redis_db")
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				beego.Info(err)
				return nil, err
			}
			return c, nil
		},
	}
}

//Subscribe 订阅基于redis的房间消息
func Subscribe(room int) redis.PubSubConn {
	psc := redis.PubSubConn{Conn: RedisConnPool.Get()}
	psc.Subscribe("chat_room_" + strconv.Itoa(room) + "_channel")
	return psc
}

//UnSubscribe 连接断开，取消订阅
func UnSubscribe(psc redis.PubSubConn) {
	psc.Unsubscribe()
	psc.Close()
}
