package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func run() (err error) {

	partitionList, err := kafkaClient.client.Partitions(kafkaClient.topic)
	if err != nil {
		logs.Error("failed to get the list of partitions:", err)
		return
	}

	for partition := range partitionList {
		pc, errRet := kafkaClient.client.ConsumePartition(kafkaClient.topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logs.Error("failed to start consumer for partition %d: %s\n", partition, err)
			return
		}

		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			kafkaClient.wg.Add(1)
			for msg := range pc.Messages() {
				logs.Debug("Parition: %d, Offset: %d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				err = sendToEs(kafkaClient.topic, msg.Value)
				if err != nil {
					logs.Warn("send to es failed,err %v", err)
				}
			}
			kafkaClient.wg.Done()
		}(pc)
	}

	kafkaClient.wg.Wait()

	return
}
