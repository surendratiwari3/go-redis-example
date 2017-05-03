//first it will subscribe to redis on one socket and it will wait for subscribe event
//second that socket will be open in a thread
//it will connect to kamailio server by provided ip
//there will be multiple kamailio so multiple client socket thread


package main

import (
"fmt" 
"github.com/go-redis/redis"
"strings"
)

func main() {
    fmt.Printf("Conneting To Redis \n")
    RedisConnect()
}

func RedisConnect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	fmt.Println("Succesfull Connected To Redis Server")
	
	//db := 0
	notificationChannel := fmt.Sprint("__keyspace@0__:kamailio:zentrunk:appserver:*")

	pubsub := client.PSubscribe(notificationChannel)
	//defer pubsub.Close()

	//err := client.Publish("mychannel1", "hello").Err()
	//if err != nil {
	//	panic(err)
	//}
    for (true) {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Channel, msg.Payload)
		//getting the key that is changes recently 
		redis_key := strings.Split(msg.Channel, "__:")
		fmt.Printf("%q\n", redis_key[1])
		args := []string{"ev_server", "ev_port"}
		ev_socket := client.HMGet(redis_key[1],args...)
		fmt.Printf("%q\n", ev_socket)
	}
}


