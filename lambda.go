package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/kelseyhightower/envconfig"
)

type Event struct {
	State      string `json:"state"` // start, stop
	InstanceID string `json:"instanceID"`
}

func HandleRequest(ctx context.Context, event Event) (string, error) {
	var awsCfg Config
	envconfig.MustProcess("AWS", &awsCfg)

	// Load session from shared config
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsCfg.Region),
		Credentials: credentials.NewStaticCredentials(awsCfg.AccessKeyID, awsCfg.SeceretKey, awsCfg.SessionToken),
	}))

	// Create new EC2 client
	svc := ec2.New(sess)

	switch event.State {
	case "start":
		err := startInstance(svc, event.InstanceID)
		if err != nil {
			return "stop error", err
		}
		return "start success", nil
	case "stop":
		err := stopInstance(svc, event.InstanceID)
		if err != nil {
			return "stop error", err
		}
		return "stop success", nil
	default:
		log.Println("state not support")
		return "error", errors.New("state not support")
	}
}

func main() {
	lambda.Start(HandleRequest)
}
