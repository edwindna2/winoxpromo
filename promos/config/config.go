package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	//KAFKA
	PROMO_BROKER_SERVERS		= "PROMO_BROKER_SERVERS"
	PROMO_BROKER_SERVERS_VAL	= "127.0.0.1:9092"
	PROMO_CLIENT_ID				= "PROMO_CLIENT_ID"
	PROMO_CLIENT_ID_VAL			= "producerPromo"
	PROMO_TOPIC             	= "PROMO_TOPIC"
	PROMO_TOPIC_VAL             = "promos"
)

func Setup() {
	//DEFAULT VALUES PULSAR
	viper.SetDefault(PROMO_BROKER_SERVERS, PROMO_BROKER_SERVERS_VAL)
	viper.SetDefault(PROMO_TOPIC, PROMO_TOPIC_VAL)

	//Get env broker servers
	brokerServers := os.Getenv(PROMO_BROKER_SERVERS)
	if len(strings.TrimSpace(brokerServers)) > 0 {
		viper.SetDefault(PROMO_BROKER_SERVERS, strings.TrimSpace(brokerServers))
	}

	//Get env topic
	topic := os.Getenv(PROMO_TOPIC)
	if len(strings.TrimSpace(topic)) > 0 {
		viper.SetDefault(PROMO_TOPIC, strings.TrimSpace(topic))
	}	
}
