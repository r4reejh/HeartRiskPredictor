package main

import (
	"fmt"
	"time"

	"github.com/gocelery/gocelery"
)

var celeryBroker *gocelery.RedisCeleryBroker
var celeryBackend *gocelery.RedisCeleryBackend
var celeryClient *gocelery.CeleryClient

func initConnector() {
	celeryBroker = gocelery.NewRedisCeleryBroker("redis://localhost:6379")
	celeryBackend = gocelery.NewRedisCeleryBackend("redis://localhost:6379")
	celeryClient, _ = gocelery.NewCeleryClient(celeryBroker, celeryBackend, 0)
	fmt.Println("connection initiated")
}

func broker(vv []float64) (label float64, isOk bool) {
	asyncResult, err := celeryClient.Delay("task.add", vv)
	if err != nil {
		panic(err)
	}

	res, err := asyncResult.Get(5 * time.Second)
	if err != nil {
		isOk = false
		label = 0
		fmt.Println("operation failed")
		return
	}

	isOk = true
	label = res.(float64)
	return
}
