package main

import "time"

// InputStruct is the struct for predict route
type InputStruct struct {
	Token    string
	Name     string
	Age      float64
	Sex      float64
	Cp       float64
	Trestbps float64
	Chol     float64
	Fbs      float64
	Restecg  float64
	Thalach  float64
	Exang    float64
	Oldpeak  float64
	Slope    float64
	Ca       float64
	Thal     float64
	UserID   int
}

type errorStruct struct {
	Message string
	Status  int
}

type loginResponse struct {
	Message string
	Status  int
	Token   string
	Name    string
}

type responseStruct struct {
	Name  string
	Label float64
	Date  time.Time
}

// Credentials  is the struct for login page
type Credentials struct {
	Username string
	Password string
}

// User type to store user in DB and parse form in signup requests
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Name     string
}

// Scan type to store the entire scan details
type Scan struct {
	ID        int
	UserID    int
	Username  string
	Name      string
	Age       float64
	Sex       float64
	Cp        float64
	Trestbps  float64
	Chol      float64
	Fbs       float64
	Restecg   float64
	Thalach   float64
	Exang     float64
	Oldpeak   float64
	Slope     float64
	Ca        float64
	Thal      float64
	Result    float64
	CreatedAt time.Time
}
