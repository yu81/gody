package gody

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
)

type Option struct {
	profile  string
	region   string
	endpoint string
}

func NewService(profile string, region string, endpoint string) (*dynamodb.DynamoDB, error) {
	svc, err := dynamodb.New(config.Config{
		Region:   region,
		Profile:  profile,
		Endpoint: endpoint,
	})
	if err != nil {
		log.Fatal("create service error")
	}
	return svc, err
}
