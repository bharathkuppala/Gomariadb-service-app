package kafkaservice

import (
	"fmt"

	utility "github.com/mariaDB/module/utilities"

	"github.com/Shopify/sarama"
)

// InitProducer ...
func InitProducer() (sarama.SyncProducer, error) {
	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{utility.KafkaConn}, config)

	return prd, err
}

// Publish ...
func Publish(message string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: utility.Topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}
}
