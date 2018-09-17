package main

import (
	"fmt"
	"time"

	"github.com/gocelery/gocelery"
)

func valueTest() {
	celeryBroker = gocelery.NewRedisCeleryBroker("redis://localhost:6379")
	celeryBackend = gocelery.NewRedisCeleryBackend("redis://localhost:6379")
	celeryClient, _ = gocelery.NewCeleryClient(celeryBroker, celeryBackend, 0)
	fmt.Println("connection initiated TEST-V")

	fmt.Print("*")
	//asyncResult, err := celeryClient.Delay("task.add", []float64{57, 0, 1, 130, 236, 0, 0, 174, 0, 0, 1, 1, 2})
	asyncResult, err := celeryClient.Delay("task.add", []float64{43, 1, 0, 120, 177, 0, 0, 120, 1, 1.2, 1, 0, 3})
	if err != nil {
		panic(err)
	}

	res, err := asyncResult.Get(2 * time.Second)
	if err != nil {
		fmt.Println()
		fmt.Println("operation failed")
		return
	}
	fmt.Println()
	fmt.Println(res)
}
