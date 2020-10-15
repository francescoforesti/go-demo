package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/francescoforesti/go-demo/goka/logging"
	"github.com/francescoforesti/go-demo/goka/utils"
	"os"
)

var (
	brokers   = initializeBrokers()
	kafkaConf = newKafkaConfiguration()
)

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

func ProducerTopic() string {
	kafkaTopic := os.Getenv(utils.KAFKA_TOPIC_STRINGS_ENV)
	if kafkaTopic == "" {
		kafkaTopic = utils.KAFKA_TOPIC_STRINGS
	}
	return kafkaTopic
}

func ConsumerTopic() string {
	kafkaTopic := os.Getenv(utils.KAFKA_TOPIC_REVERSED_ENV)
	if kafkaTopic == "" {
		kafkaTopic = utils.KAFKA_TOPIC_REVERSED
	}
	return kafkaTopic
}

func CreateKafkaProducer() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(brokers, kafkaConf)
	//defer producer.Close()
	if err != nil {
		logging.Error(
			fmt.Sprintf("Kafka error in Producer Initialization: %s", err))
		os.Exit(1)
	}
	logging.Info("Kafka Producer initialized")
	return producer, nil
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

func CreateKafkaConsumer() (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer(brokers, newKafkaConfiguration())
	if err != nil {
		logging.Error(fmt.Sprintf("Kafka Consumer creation error: %s", err))
		os.Exit(1)
	}
	return consumer, nil
}
