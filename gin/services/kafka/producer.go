package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/francescoforesti/go-demo/gin/logging"
	"github.com/francescoforesti/go-demo/gin/models"
	"time"
)

type MessageEvent struct {
	Content   *models.GinMessage
	Timestamp *time.Time
}

func ProduceEvent(producer *sarama.SyncProducer, content *models.GinMessage) error {
	event := createEvent(content)
	topic := ProducerTopic()
	return sendMsg(*producer, topic, *event)
}

func sendMsg(producer sarama.SyncProducer, topic string, event MessageEvent) error {
	logging.Info(fmt.Sprintf("Message: %+v", event))
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(json),
	}

	partition, offset, err := producer.SendMessage(msgLog)
	if err != nil {
		logging.Error(fmt.Sprintf("Kafka error in Sending Message: %s", err))
	} else {
		logging.Info(fmt.Sprintf("Message is stored in partition %d, offset %d", partition, offset))
	}
	return err
}

func createEvent(content *models.GinMessage) *MessageEvent {
	event := new(MessageEvent)
	event.Content = content
	event.Timestamp = createT(time.Now())
	return event
}

func createT(t time.Time) *time.Time {
	return &t
}
