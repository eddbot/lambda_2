package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	EmployeeId  string `json:"employeeId"`
	LeavingDate string `json:"leavingDate"`
	Email       string `json:"email"`
}

type Output struct {
	EmployeeId    string `json:"employeeId"`
	LeavingDate   string `json:"leavingDate"`
	Email         string `json:"email"`
	ExecutionTime int    `json:"executionTime"`
}

func handler(ctx context.Context, i Input) (Output, error) {

	// this could be updated to use a dynamic timestamp path
	// rather than the seconds directly
	// https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-wait-state.html

	future := time.Now().Add(time.Hour * (180 * 24))
	now := time.Now()
	timeToWaitInSeconds := int(future.Sub(now).Seconds())

	out := Output{
		EmployeeId:  i.EmployeeId,
		LeavingDate: i.LeavingDate,
		Email:       i.Email,
		// this employee is leaving in 180 days
		ExecutionTime: timeToWaitInSeconds,
	}
	return out, nil
}
func main() {
	lambda.Start(handler)
}
