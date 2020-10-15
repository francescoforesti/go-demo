package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/francescoforesti/go-demo/kafka/logging"
	"os"
)

func CreateKafkaConsumer() (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer(brokers, newKafkaConfiguration())
	if err != nil {
		logging.Error(fmt.Sprintf("Kafka Consumer creation error: %s", err))
		os.Exit(1)
	}
	return consumer, nil
}

func initializeConsumer(topic string) (sarama.PartitionConsumer, error) {
	kafkaConsumer, err := CreateKafkaConsumer()
	var partitionConsumer sarama.PartitionConsumer
	if kafkaConsumer != nil {
		partitions, err := kafkaConsumer.Partitions(topic)
		partitionConsumer, err = kafkaConsumer.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if err != nil {
			logging.Error(
				fmt.Sprintf("Error during Kafka Consumer initialization for topic %s: %s", topic, err))
			os.Exit(1)
		}
		logging.Info("Kafka consumer initialized")
	} else if err != nil {
		logging.Error("Error during Kafka partition initialization: " + err.Error())
	}
	return partitionConsumer, err
}

func ConsumeEvents(partitionConsumer sarama.PartitionConsumer) ([]byte, error) {
	var msgVal []byte
	var err error

	fmt.Printf("Waiting for messages...\n\n")
	select {
	case err := <-partitionConsumer.Errors():
		logging.Error(fmt.Sprintf("Kafka error in Consuming Messages: %s", err))
	case msg := <-partitionConsumer.Messages():
		msgVal = msg.Value
	}
	return msgVal, err
}
