package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	testBroker()
	initConnector()
	http.HandleFunc("/predict", predict)
	http.HandleFunc("/test", test)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		Addr:         ":8082",
	}
	log.Fatal(srv.ListenAndServe())
}

func predict(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t inputStruct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	inputValues := []float64{t.age, t.sex, t.cp, t.trestbps, t.chol, t.fbs, t.restecg, t.thalach, t.exang, t.oldpeak, t.slope, t.ca, t.thal}
	x, isOk := broker(inputValues)
	if isOk != true {
		errorX := errorStruct{}
		errorX.Message = "Some Error"
		errorX.Status = 500

		erb, err := json.Marshal(errorX)
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(200)
		rw.Write(erb)
	} else {
		response := responseStruct{}
		response.Label = x

		rb, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(200)
		rw.Write(rb)
	}
}

func test(rw http.ResponseWriter, req *http.Request) {
	response := responseStruct{}
	response.Label = 1

	rb, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(200)
	rw.Write(rb)
}
