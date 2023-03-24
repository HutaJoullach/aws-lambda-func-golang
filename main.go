package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"How old are you"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	cdate := time.Now()
	cy, _, _ := cdate.Date()
	yob := cy - event.Age

	var gen string
	if yob >= 1926 && yob <= 1945 {
		gen = "Silent Generation"
	} else if yob >= 1946 && yob <= 1965 {
		gen = "Baby Boomers"
	} else if yob >= 1966 && yob <= 1979 {
		gen = "Generation X"
	} else if yob >= 1980 && yob <= 1995 {
		gen = "Millennials"
	} else if yob >= 1996 {
		gen = "Generation Z"
	} else {
		gen = "not in the group category"
	}

	return MyResponse{Message: fmt.Sprintf("Hey %s 👋 , and you are %s", event.Name, gen)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
