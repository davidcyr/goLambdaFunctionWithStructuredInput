package main

// Simple Lambda Function
//   that receives a structured JSON request
//   and replies with a structured JSON response


import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string "json:name"
	Age  int    "json:age"
}

// TODO - investigate
// BIT OF A MYSTERY HERE -- 
// why is the json:code / json:Answer mapping not working for output??
type MyResponse struct {
	Code string "json:code"
	Message string "json:Answer"
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Code: "200", Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
