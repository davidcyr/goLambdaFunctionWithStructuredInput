package main

// Simple Lambda Function
//   that receives a structured JSON request
//   and replies with a structured JSON response


import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

}

// see https://golang.org/pkg/encoding/json/#Marshal for more information on formatting options
// this shows an example of omitting a field if it is empty
// puls a 'control group' of an EmptyField to show default handling
type MyResponse struct {
	Code string `json:"code,omitempty"`
	Message string `json:"answer"`
	EmptyField string `json:"emptyfield"`
}


func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	if (event.Age == 50) {
		return MyResponse{Code: "200", Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
	}
	return MyResponse{ Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
