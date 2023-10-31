package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var (
	hostQueueName   string
	requestMessage  string
	responseMessage string
)

func init() {
	flag.StringVar(&hostQueueName, "host-queue", "request-queue", "Name of the Host Request Queue")
	flag.StringVar(&requestMessage, "request-message", "Example Request Message", "An Example Message to send to the Virtual Queue")
	flag.StringVar(&responseMessage, "response-message", "Example Response Message", "The message that will be responded with")
}

func main() {
	flag.Parse()
	// Initialize config that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println(err)
		return
	}

	sqsClient := sqs.NewFromConfig(cfg)

	// Create a queue to listen on
	log.Println("creating host queue:", hostQueueName)
	hostqueueout, err := sqsClient.CreateQueue(context.Background(), &sqs.CreateQueueInput{
		QueueName:  aws.String(hostQueueName),
		Attributes: map[string]string{},
	})
	if err != nil {
		log.Println(err)
		return
	}

	// Delete the queue when we're done
	defer func() {
		log.Println("deleting host queue:", hostQueueName)
		_, err = sqsClient.DeleteQueue(context.Background(), &sqs.DeleteQueueInput{
			QueueUrl: hostqueueout.QueueUrl,
		})
		if err != nil {
			log.Println(err)
			return
		}
	}()
}
