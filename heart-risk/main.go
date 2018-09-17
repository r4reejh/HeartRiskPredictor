package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	valueTest()
	testBroker()
	initConnector()

	http.HandleFunc("/predict", predict)
	http.HandleFunc("/test", test)
	http.HandleFunc("/login", login)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		Addr:         ":8082",
	}
	log.Fatal(srv.ListenAndServe())
}

func predict(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var T InputStruct
	err := decoder.Decode(&T)
	if err != nil {
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
		panic(err)
	}

	inputValues := []float64{T.Age, T.Sex, T.Cp, T.Trestbps, T.Chol, T.Fbs, T.Restecg, T.Thalach, T.Exang, T.Oldpeak, T.Slope, T.Ca, T.Thal}
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
		response.Name = T.Name
		response.Date = time.Now()

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
