package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/francescoforesti/go-demo/goka/models"
	"github.com/francescoforesti/go-demo/kafka/logging"
	"github.com/francescoforesti/go-demo/kafka/utils"
	"os"
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
		content, err := ConsumeEvents(consumer)
		if err != nil {
			var msg models.Strings
			msg.Message = string(reverse(content))
			err = ProduceEvent(&producer, &msg)
			if err != nil {
				logging.Warn("Error producing event")
			}
		} else {
			logging.Warn("Error consuming event")
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
	//conf.ChannelBufferSize = 1
	conf.Version = sarama.V2_6_0_0
	return conf
}

func ConsumerTopic() string {
	kafkaTopic := os.Getenv(utils.KAFKA_TOPIC_STRINGS_ENV)
	if kafkaTopic == "" {
		kafkaTopic = utils.KAFKA_TOPIC_STRINGS
	}
	return kafkaTopic
}

func ProducerTopic() string {
	kafkaTopic := os.Getenv(utils.KAFKA_TOPIC_REVERSED_ENV)
	if kafkaTopic == "" {
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
	// TODO: completare
	return *req
}
