package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {

	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		fmt.Printf("failed to start consumer: %s", err)
		return
	}

	partitionList, err := consumer.Partitions("test1")
	if err != nil {
		fmt.Println("failed to get the list of partitions:", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test1", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d: %s\n", partition, err)
			return
		}

		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			wg.Add(1)
			for msg := range pc.Messages() {
				fmt.Printf("Parition: %d, Offset: %d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
			wg.Done()
		}(pc)
	}

	//time.Sleep(time.Hour)
	fmt.Println("wait....")
	wg.Wait()
	_ = consumer.Close()
	fmt.Println("over...")

}
