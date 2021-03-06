package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	kafkaModels "github.com/francescoforesti/go-demo/kafka/models"
	"github.com/francescoforesti/go-demo/logging"
	"os"
)

func CreateKafkaProducer() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(brokers, kafkaConf)
	if err != nil {
		logging.Error(
			fmt.Sprintf("Kafka error in Producer Initialization: %s", err))
		os.Exit(1)
	}
	logging.Info("Kafka Producer initialized")
	return producer, nil
}

func ProduceEvent(producer *sarama.SyncProducer, content *kafkaModels.KafkaMessage) error {
	topic := ProducerTopic()
	return sendMsg(*producer, topic, *content)
}

func sendMsg(producer sarama.SyncProducer, topic string, msg kafkaModels.KafkaMessage) error {
	logging.Info(fmt.Sprintf("Message: %+v", msg))
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonMsg),
	}

	partition, offset, err := producer.SendMessage(msgLog)
	if err != nil {
		logging.Error(fmt.Sprintf("Kafka error in Sending Message: %s", err))
	} else {
		logging.Info(fmt.Sprintf("Message is stored in partition %d, offset %d", partition, offset))
	}
	return err
}
