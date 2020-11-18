package handlers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/francescoforesti/go-demo/gin/models"
	"github.com/francescoforesti/go-demo/gin/services/kafka"
	"github.com/francescoforesti/go-demo/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	producer    sarama.SyncProducer
	errProducer error
	consumer    sarama.PartitionConsumer
	errConsumer error
)

func InitializeHandlers() {
	kafka.CreateTopic(kafka.ProducerTopic())
	kafka.CreateTopic(kafka.ConsumerTopic())
	producer, errProducer = kafka.CreateKafkaProducer()
	consumer, errConsumer = initializeConsumer(kafka.ConsumerTopic())
}

func initializeConsumer(topic string) (sarama.PartitionConsumer, error) {
	kafkaConsumer, err := kafka.CreateKafkaConsumer()
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

func PostString(c *gin.Context) {
	logging.Debug(c.Request.Method + " " + c.Request.URL.String())
	var model *models.GinMessage
	errBindJSON := c.BindJSON(&model)
	if errBindJSON != nil {
		logging.Warn("errBindJSON not nil: " + errBindJSON.Error())
		return
	}
	if errProducer != nil {
		logging.Warn("errProducer not nil: " + errProducer.Error())
		return
	}
	logging.Info(model.Message)
	err := kafka.ProduceEvent(&producer, model)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			},
		)
	}
}

func GetReversedString(c *gin.Context) {
	logging.Debug(c.Request.Method + " " + c.Request.URL.String())
	msg := *<-consumer.Messages()
	c.Status(http.StatusOK)
	c.Header("Content-Type", "application/json")
	_, _ = c.Writer.Write(msg.Value)
}
