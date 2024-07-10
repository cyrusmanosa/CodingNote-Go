package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	Address  = "localhost:6379"
	Password = ""
	DB       = 1
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
	SetMapValueInDB("jsonTest3", data, client)

	result, err := GetMapValueInDB("jsonTest3", client)
	chcekPanic(err)
	fmt.Println(result)
}

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     Address,
		Password: Password,
		// DB:       DB,
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
