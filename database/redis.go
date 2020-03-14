package database

import (
	"bug-management/conf"
	. "common/logs"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisSingle *redis.Client

func init(){
	fmt.Printf("Redis run in single mode\n")
	redisSingle = redis.NewClient(&redis.Options{
		Addr:     conf.RsAddr,
		Password: conf.RsPwd, // "" : no password set
		DB:       conf.Rsdb,  // 0  : use default DB
	})

	err := redisSingle.Ping().Err()
	if err != nil {
		fmt.Printf("Connect to redis failed:%v \n", err)
		panic("Connect to redis failed")
	}
}


func RsSet(rskey string, rsvalue string, rsexpire int) (err error) {
	err = redisSingle.Set(rskey, rsvalue, time.Duration(rsexpire)*time.Second).Err()
	if err != nil {
		//	panic(err)
		Error("redis. Set key to error", err)
		return err
	}
	return nil
}

func RsGet(rskey string) (err error, val string) {
	val, err = redisSingle.Get(rskey).Result()
	if err != nil {
		//	panic(err)
		Error("redis. Get key from error", err)
		return err, ""
	}
	return nil, val
}

func  RsDel(rskey string) (err error) {
	err = redisSingle.Del(rskey).Err()
	if err != nil {
		//	panic(err)
		Error("redis. Del key from error", err)
		return err
	}
	Info("RsDel: key=[", rskey, "].")
	return nil
}
