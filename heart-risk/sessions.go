package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

var cache redis.Conn

func initSessionCache() {
	conn, err := redis.DialURL("redis://localhost:6379")
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
		v, err := cache.Do("SETEX", sessionToken, "120", creds.Username)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}
		fmt.Println(v)

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

func checkSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		c, err := req.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				forbidden := errorStruct{}
				forbidden.Message = "You are not logged in"
				forbidden.Status = 401

				frb, err := json.Marshal(forbidden)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
				}
				rw.Header().Set("content-type", "application/json")
				rw.WriteHeader(200)
				rw.Write(frb)
				return
			}
			// For any other type of error, return a bad request status
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value
		response, err := cache.Do("GET", sessionToken)
		if err != nil {
			// If there is an error fetching from cache, return an internal server error status
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if response == nil {
			forbidden := errorStruct{}
			forbidden.Message = "You are not logged in"
			forbidden.Status = 401

			frb, err := json.Marshal(forbidden)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			rw.Header().Set("content-type", "application/json")
			rw.WriteHeader(200)
			rw.Write(frb)
			return
		}
		next.ServeHTTP(rw, req)
	})
}
