package source

import (
	"fmt"

	config "gihub.com/dna2/promos/config"
	model "gihub.com/dna2/promos/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

type kafkaSrc struct {
	producer	*kafka.Producer
	topic		*string
}

func newInstanceBrokerSrc() (*kafkaSrc, error) {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": viper.GetString(config.PROMO_BROKER_SERVERS),
		"client.id": viper.GetString(config.PROMO_CLIENT_ID),
		"acks": "all",
		"retries":2,
	}
	producer, err := kafka.NewProducer(&configMap)
	if err != nil {
		return nil, fmt.Errorf("Kafka: error producer - %v", err)
	}

	topic := viper.GetString(config.PROMO_TOPIC)
	return &kafkaSrc{producer: producer, topic: &topic}, nil
}

func (kafkaSrc *kafkaSrc) SendMsg(msg model.MessageBase)(*kafka.Message,error){

	msgSerialize,err := msg.ToBytes()
	if err != nil {
		return nil, fmt.Errorf("Kafka: error sendMsg - %v", err)
	}

	delivery_chan := make(chan kafka.Event)
	err = kafkaSrc.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: kafkaSrc.topic, Partition: kafka.PartitionAny},
		Value: msgSerialize,
	},
		delivery_chan,
	)

	if err!= nil {
		return nil, fmt.Errorf("Kafka: error producer - %v", err)
	}

	e := <-delivery_chan
	m := e.(*kafka.Message)
	
	if m.TopicPartition.Error != nil {
		return nil, fmt.Errorf("Kafka: Delivery failed - %v", m.TopicPartition.Error)
	}

	//log.Printf("Kafka: Delivered message to topic %s [%d] at offset %v\n",*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	return m, nil
}

func (kafkaSrc *kafkaSrc) close() {
	if kafkaSrc.producer != nil {
		kafkaSrc.producer.Flush(2000)
	}
}