package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err :", err)
		return
	}

	defer conn.Close()
	redisKey(conn)

	redisHash(conn)
}

func redisKey(conn redis.Conn) {

	_, err := conn.Do("Set", "name", "田利军")
	if err != nil {
		fmt.Println("conn.Do :", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.DO", err)
		return
	}

	fmt.Println("操作完成 ok", r)
}

func redisHash(conn redis.Conn) {

	// 逐一存放   Hset user01 key value
	// _, err := conn.Do("Hset", "user01", "name", "田利军")
	// if err != nil {
	// 	fmt.Println("conn.Do :", err)
	// 	return
	// }

	// _, err = conn.Do("Hset", "user01", "sex", "男")
	// if err != nil {
	// 	fmt.Println("conn.Do :", err)
	// 	return
	// }

	// _, err = conn.Do("Hset", "user01", "age", "25")
	// if err != nil {
	// 	fmt.Println("conn.Do :", err)
	// 	return
	// }

	// r, err := redis.String(conn.Do("Hget", "user01", "name"))
	// if err != nil {
	// 	fmt.Println("conn.DO", err)
	// 	return
	// }

	// fmt.Println("user01 name:", r)

	// r, err = redis.String(conn.Do("Hget", "user01", "age"))
	// if err != nil {
	// 	fmt.Println("conn.DO", err)
	// 	return
	// }

	// fmt.Println("user01 age:", r)

	// r, err = redis.String(conn.Do("Hget", "user01", "sex"))
	// if err != nil {
	// 	fmt.Println("conn.DO", err)
	// 	return
	// }

	// fmt.Println("user01 sex:", r)

	//---------------------------------------------------------------------------------

	//批量存放和读取  Hmset user01 key1 value1 key2 value2 key value3

	_, err := conn.Do("HMSet", "user02", "name", "jiashaofen", "sex", "男", "age", "26")
	if err != nil {
		fmt.Println("conn.DO", err)
		return
	}

	fmt.Println("切割")
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "sex", "age"))
	if err != nil {
		fmt.Println("conn.DO", err)
		return
	}

	for i, value := range r {

		fmt.Printf("r[%d]=%v\n", i, value)
	}
	fmt.Println("操作完成 ok")
}
