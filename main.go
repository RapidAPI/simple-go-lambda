package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

var (
	logger = log.Default()
)

func main() {
	lambda.Start(Handler)
}
func Handler(_ context.Context) {
	defer func() {
		if r := recover(); r != nil {
			print("recovered again", "\n")
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			logger.Panic("recovered from ", r)
		}
	}()
	fmt.Println("DEBUG", "fmt.Println")
	fmt.Println("WARN", "fmt.Println")
	fmt.Println("ERROR", "fmt.Println")
	fmt.Printf("DEBUG %s\n", "fmt.Printf")
	fmt.Printf("WARN %s", "fmt.Printf\n")
	fmt.Printf("ERROR %s", "fmt.Printf\n")
	println("DEBUG", "println")
	println("WARN", "println")
	println("ERROR", "println")
	logger.Println("DEBUG", "logger.Println")
	logger.Printf("DEBUG:%s", "logger.Printf")
	logger.Fatalln("logger.Fatalln")
}
