package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

var cache redis.Conn

func initSessionCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}

func login(rw http.ResponseWriter, req *http.Request) {
	creds := Credentials{}
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		panic(err)
	}

	sessionToken, err := uuid.NewV4()
	if creds.Username == "r4reejh" && creds.Password == "test" {
		http.SetCookie(rw, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken.String(),
			Expires: time.Now().Add(120 * time.Second),
		})

		successMessage := errorStruct{}
		successMessage.Message = "Logged in successfully"
		successMessage.Status = 200

		srb, err := json.Marshal(successMessage)
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(200)
		rw.Write(srb)
	} else {
		successMessage := errorStruct{}
		successMessage.Message = "Credentials Incorrect"
		successMessage.Status = 400

		srb, err := json.Marshal(successMessage)
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(200)
		rw.Write(srb)
	}
}
