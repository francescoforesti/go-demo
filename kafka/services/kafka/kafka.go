package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/francescoforesti/go-demo/gin/models"
	kafkaModel "github.com/francescoforesti/go-demo/kafka/models"
	"github.com/francescoforesti/go-demo/kafka/utils"
	"github.com/francescoforesti/go-demo/logging"
	"os"
	"time"
)

var (
	brokers   = initializeBrokers()
	kafkaConf = newKafkaConfiguration()
)
var producer sarama.SyncProducer
var errProducer error
var consumer sarama.PartitionConsumer
var errConsumer error

func InitializeHandlers() {
	CreateTopic(ProducerTopic())
	CreateTopic(ConsumerTopic())
	producer, errProducer = CreateKafkaProducer()
	consumer, errConsumer = initializeConsumer(ConsumerTopic())

	logging.Info("starting to pipe messages")
	startPiping()
}

func startPiping() {
	for {
		content, _ := ConsumeEvents(consumer)
		var kafkaMsg kafkaModel.KafkaMessage
		err := json.Unmarshal(content, &kafkaMsg)
		if err != nil {
			logging.Warn("Error unmarshalling event")
		}
		var ginMessage = kafkaMsg.Content
		kafkaMsg.Reversed = models.GinMessage{Message: string(reverse([]byte(ginMessage.Message)))}
		var now = time.Now()
		kafkaMsg.Timestamp = &now
		err2 := ProduceEvent(&producer, &kafkaMsg)
		if err2 != nil {
			logging.Warn("Error producing event")
		}
	}
}

func reverse(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func initializeBrokers() []string {
	broker := os.Getenv(utils.KAFKA_BROKER_ADDRESS_ENV)
	if len(broker) == 0 {
		broker = utils.KAFKA_BROKER_ADDRESS
	}
	return []string{broker}
}

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_6_0_0
	return conf
}

func ConsumerTopic() string {
	kafkaTopic, available := os.LookupEnv(utils.KAFKA_TOPIC_STRINGS_ENV)
	if !available {
		kafkaTopic = utils.KAFKA_TOPIC_STRINGS
	}
	return kafkaTopic
}

func ProducerTopic() string {
	kafkaTopic, available := os.LookupEnv(utils.KAFKA_TOPIC_REVERSED_ENV)
	if !available {
		kafkaTopic = utils.KAFKA_TOPIC_REVERSED
	}
	return kafkaTopic
}

func CreateTopic(topic string) {
	kafkaBroker := sarama.NewBroker(brokers[0])
	err := kafkaBroker.Open(kafkaConf)
	if err != nil {
		logging.Error(fmt.Sprintf("Cannot initialize topic '%s': %s\n", topic, err.Error()))
	}
	request := newCreateTopicRequest()
	topics, err := kafkaBroker.CreateTopics(&request)
	if err != nil {
		logging.Error(fmt.Sprintf("Cannot initialize topic '%s': %s\n", topic, err.Error()))
	} else {
		logging.Info(fmt.Sprintf("topic %s created: %d\n", topic, topics.Version))
	}
}

func newCreateTopicRequest() sarama.CreateTopicsRequest {
	req := new(sarama.CreateTopicsRequest)
	// TODO
	return *req
}
