package main

import (
	"flag"
	"log"
)

// ServiceConfig defines all of the service configuration parameters
type ServiceConfig struct {
	QueueName         string
	MessageBucketName string
}

// LoadConfiguration will load the service configuration from env/cmdline
// and return a pointer to it. Any failures are fatal.
func LoadConfiguration() *ServiceConfig {

	var cfg ServiceConfig
	flag.StringVar(&cfg.QueueName, "queue", "", "Queue name")
	flag.StringVar(&cfg.MessageBucketName, "bucket", "", "Oversize message bucket name")

	flag.Parse()

	if len(cfg.QueueName) == 0 {
		log.Fatalf("QueueName cannot be blank")
	}
	if len(cfg.MessageBucketName) == 0 {
		log.Fatalf("MessageBucketName cannot be blank")
	}

	log.Printf("[CONFIG] QueueName         = [%s]", cfg.QueueName)
	log.Printf("[CONFIG] MessageBucketName = [%s]", cfg.MessageBucketName)

	return &cfg
}

//
// end of file
//
