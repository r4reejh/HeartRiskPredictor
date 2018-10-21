package main

import (
	"encoding/json"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func info(rw http.ResponseWriter, req *http.Request) {
	var T map[string]string
	err := json.NewDecoder(req.Body).Decode(&T)
	checkErr(err)

	if c, ok := req.Header["X-Hp-Token"]; ok {
		username, err := redis.String(cache.Do("GET", c[0]))
		checkErr(err)
		history := DBfetchUserHistory(username)

		rb, err := json.Marshal(history)
		checkErr(err)

		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		rw.Header().Set("Content-type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.WriteHeader(200)
		rw.Write(rb)
	} else {
		errorMsg := errorStruct{}
		errorMsg.Message = "token invalid"
		errorMsg.Status = 400

		rb, err := json.Marshal(errorMsg)
		checkErr(err)

		rw.WriteHeader(200)
		rw.Write(rb)
	}
}
