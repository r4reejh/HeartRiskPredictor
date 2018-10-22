package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	valueTest()
	//testBroker()
	initSessionCache()
	initConnector()
	initRegex()
	dbInit()

	predictHandler := http.HandlerFunc(predict)
	loginHandler := http.HandlerFunc(login)
	signupHandler := http.HandlerFunc(signup)
	testHandler := http.HandlerFunc(test)
	infoHandler := http.HandlerFunc(info)

	http.Handle("/predict", corsHandler(checkSessionToken(predictHandler)))
	http.Handle("/test", checkSessionToken(testHandler))
	http.Handle("/login", corsHandler(loginHandler))
	http.Handle("/signup", corsHandler(signupHandler))
	http.Handle("/info", corsHandler(infoHandler))
	http.HandleFunc("/echo", echo)

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
		return
	}
	response := responseStruct{}
	response.Label = x
	response.Name = T.Name
	response.Date = time.Now()

	defer DBrecordScan(T, x)
	rb, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(200)
	rw.Write(rb)
	return
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

func echo(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var T InputStruct
	err := decoder.Decode(&T)
	erb, err := json.Marshal(T)
	if err != nil {
		panic(err)
	}

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token, X-Hp-Token"
	rw.Header().Set("Content-type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	rw.WriteHeader(200)
	rw.Write(erb)
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token, X-Hp-Token"
		w.Header().Set("Content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			//handle preflight in here
			w.WriteHeader(200)
			return
		}
		h.ServeHTTP(w, r)
	}
}
