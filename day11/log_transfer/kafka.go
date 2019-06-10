package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
)

var (
	// wg sync.WaitGroup
	kafkaClient *KafkaClient
)

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
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
		partitionList, err := consumer.Partitions(topic)
		if err != nil {
			logs.Error("failed to get the list of partitions:", err)
			return
		}
		// fmt.Println(partitionList)

		for partition := range partitionList {
			pc, errRet := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
			if errRet != nil {
				err = errRet
				logs.Error("failed to start consumer for partition %d: %s\n", partition, err)
				return
			}

			defer pc.AsyncClose()
			go func(sarama.PartitionConsumer) {
				//wg.Add(1)
				for msg := range pc.Messages() {
					logs.Debug("Parition: %d, Offset: %d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				}
				//wg.Done()
			}(pc)
		}

		//wg.Wait()
		//_ = consumer.Close()

		return

	*/
}
