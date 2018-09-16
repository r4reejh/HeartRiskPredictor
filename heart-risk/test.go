package main

import (
	"fmt"
	"time"

	"github.com/gocelery/gocelery"
)

func testBroker() {
	celeryBroker = gocelery.NewRedisCeleryBroker("redis://localhost:6379")
	celeryBackend = gocelery.NewRedisCeleryBackend("redis://localhost:6379")
	celeryClient, _ = gocelery.NewCeleryClient(celeryBroker, celeryBackend, 0)
	fmt.Println("connection initiated TEST")

	for i := 0; i < 10; i++ {
		fmt.Print("*")
		asyncResult, err := celeryClient.Delay("task.add", []float64{57, 0, 1, 130, 236, 0, 0, 174, 0, 0, 1, 1, 2})
		if err != nil {
			panic(err)
		}

		_, err = asyncResult.Get(2 * time.Second)
		if err != nil {
			fmt.Println()
			fmt.Println("operation failed")
			return
		}
	}
	fmt.Println()
	fmt.Println("tests passed")
}
