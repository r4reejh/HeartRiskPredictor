package main

import "time"

// InputStruct is the struct for predict route
type InputStruct struct {
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
}

type errorStruct struct {
	Message string
	Status  int
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
