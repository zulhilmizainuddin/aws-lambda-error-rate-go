package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() error {
	return errors.New("something went wrong")
}

func main() {
	lambda.Start(Handler)
}
