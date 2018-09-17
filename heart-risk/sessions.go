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
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
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
	} else {
		successMessage := errorStruct{}
		successMessage.Message = "Logged in successfully"
		successMessage.Status = 200

		srb, err := json.Marshal(successMessage)
		if err != nil {
			panic(err)
		}

		rw.WriteHeader(200)
		rw.Write(srb)
	}
}
