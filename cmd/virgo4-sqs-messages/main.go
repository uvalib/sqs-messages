package main

import (
	"github.com/uvalib/virgo4-sqs-sdk/awssqs"
	"log"
)

//
// main entry point
//
func main() {

	//log.Printf("===> %s service staring up <===", os.Args[0])

	// Get config params and use them to init service context. Any issues are fatal
	cfg := LoadConfiguration()

	// load our AWS_SQS helper object
	aws, err := awssqs.NewAwsSqs(awssqs.AwsSqsConfig{MessageBucketName: cfg.MessageBucketName})
	if err != nil {
		log.Fatal(err)
	}

	// get the message count from the queue
	count, err := aws.GetMessagesAvailable(cfg.QueueName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: %d message(s)", cfg.QueueName, count)
}

//
// end of file
//
