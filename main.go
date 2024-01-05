package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func PeriodicRun(task func(...any), interval int, params ...any) {
	for {
		fmt.Println("run")
		task(params...)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func defaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  http.StatusOK,
	})
}

func statusGET(c *gin.Context) {
	ret, err := rdb.Get(context.TODO(), "takara2314_status").Result()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": ret,
		"status":  http.StatusOK,
	})
}

func choice(s []string) string {
	return s[rand.Intn(len(s))]
}

func updateStatus(params ...any) {
	if len(params) != 1 {
		panic("invalid params")
	}

	ctx := params[0].(*context.Context)

	luck := choice([]string{"good", "bad", "normal"})
	fmt.Println(luck)

	rdb.Set(*ctx, "takara2314_status", luck, 0)
}

func main() {
	fmt.Println("Hello World")

	ctx := context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis_study:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Redis connected")
	}

	fmt.Println("setup")
	updateStatus(&ctx)
	fmt.Println("periodic")
	go PeriodicRun(updateStatus, 5, &ctx)

	router := gin.Default()
	router.GET("/", defaultHandler)
	router.GET("/ping", defaultHandler)
	router.GET("/status", statusGET)
	router.Run(":8080")
}
