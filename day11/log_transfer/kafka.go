package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
	"sync"
)

var (
	kafkaClient *KafkaClient
)

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
	wg     sync.WaitGroup
}

func initKafka(addr string, topic string) (err error) {

	kafkaClient = &KafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logs.Error("failed to start kafka consumer: %s", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.addr = addr
	kafkaClient.topic = topic

	return

	/*


		//wg.Wait()
		//_ = consumer.Close()

		return

	*/
}
