/*
----------------------------------------------------------------

	SET Name Cyrus 'Name' => 'Cyrus'
	DEL Name
	GET Name

	if Set 2 time , will be reset
	example:
		name => cyrus , Set name man
		Get name => man

	[-NX] make sure to check is not Exist this item
	if name => cyrus , Set name man NX, result => nil

	[-XX] make sure to check have this item
	if have not this item => nil

----------------------------------------------------------------

	[SET GET]
	{New item}
	SET nameTest Woody GET => result: nil
	SET nameTest Cyrus GET => result: Woody
	SET nameTest Jason GET => result: Cyrus
	GET nameTest => result: Jason

	[Set multiple]
	MSET Name2 Man ago 30 address osaka
	MGET Name2 ago address

	[GETRANGE]
	GETRANGE name 0 3 => result: cyru
	GETRANGE name -3 -1 => result: rus

	[SETRANGE]
	SETRANGE name 2 1111 => result: Cy1111

	----------------------------------------------------------------

	[INCR]
	INCR ago => +1

	[DECR]
	DECR ago => -1

	[INCRBY]
	INCRBY ago 5 => +5

	[DECRBY]
	DECRBY ago 5 => -5

	----------------------------------------------------------------

	[SADD] <= add array
	SADD names Cyrus Woody Man => result: [Cyrus,Woody,Man]

	[SREM] <= remove array item
	SREM names man

	[SUNION] * only display
	SUNION array1 array2

	[SISMEMBER] * SELECT data to　confirmation
	SISMEMBER names cyrus

----------------------------------------------------------------

	[LPUSH] <= list(index,Element)
	LPUSH order aaa bbb => result: 1:aaa, 2:bbb
	LPUSH order ccc => result: 1:ccc,2:aaa,3:bbb

	[RPUSH]
	RPUSH order ddd eee => result: 1:ccc, 2:aaa, 3:bbb, 4:ddd, 5:eee

	[LPOP] <= output from first item, 1 <= number of items
	LPOP order 1 => aaa, bbb, ddd, eee

	[RPOP] <= output from last item
	RPOP order 1 => aaa, bbb, ddd

	[LINDEX] <= Select of items number
	LINDEX order 1 => result: bbb

	[LPOS] <= Select of data
	LPOS order bbb => result: 1


	[LLEN] <= len of list
	LLEN order => 3

	[LRANGE]
	LRANGE order 0 1 => result: aaa, bbb

----------------------------------------------------------------

	[HSET] <= Field , Value (map)
	HSET books:1 title "Name of the book" => (New Item:1) title : Name of the book
	HSET books:2 title "asdasd" => (New Item book:2)

	[HGETALL]
	HGETALL books:1 => Field,Value,Field,Value

	[HGET]
	HGET books:1 title => result: Name of the book

	[HDEL]
	HGET books:1 title

	[HKEYS] <= show all key
	HKEYS books:1

	[HVALS] <= show all value
	HVALS books:1


	[HEXISTS]
	HEXISTS books:1 Name of the book => result: 1

----------------------------------------------------------------

	＊＊＊  if have not item => will be Create
*/
package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	Address  = "localhost:6379"
	Password = ""
	DB       = 0
)

func main() {

	// ------------------ Connect ---------------------
	client := ConnectRedis()

	// ------------------ SetOneValue ---------------------
	// setup := SetOneValueInDB("nameTest2", "Cyrus2", client)
	// chcekPanic(setup)
	// if setup == nil {
	// 	fmt.Println("OK")
	// }

	// ------------------ GetOneValue ---------------------
	// val, err := GetValueInDB("nameTest2", client)
	// chcekPanic(err)
	// fmt.Println(val)

	// ------------------ SetMap ---------------------
	data := map[string]string{
		"Name":       "CyrusMan",
		"Age":        "30",
		"Occupation": "Student",
	}
	SetMapValueInDB("jsonTest2", data, client)

	result, err := GetMapValueInDB("jsonTest2", client)
	chcekPanic(err)
	fmt.Println(result)
}

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     Address,
		Password: Password,
		DB:       DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	return client
}

func SetOneValueInDB(key, value string, client *redis.Client) error {
	err := client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Println("Failed to set nameTest: ", err)
		return err
	}
	return err
}
func GetValueInDB(key string, client *redis.Client) (string, error) {
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("Failed to get nameTest: ", err)
		return "", err
	}
	return val, nil
}

func SetMapValueInDB(key string, value map[string]string, client *redis.Client) error {
	for k, v := range value {
		err := client.HSetNX(context.Background(), key, k, v).Err()
		if err != nil {
			fmt.Println("Failed to set nameTest: ", err)
			return err
		}
	}
	return nil
}
func GetMapValueInDB(key string, client *redis.Client) (map[string]string, error) {
	val, err := client.HGetAll(context.Background(), key).Result()
	if err != nil {
		fmt.Println("Failed to get nameTest: ", err)
		return val, err
	}
	return val, nil
}
func chcekPanic(err error) {
	if err != nil {
		panic(err)
	}
}
