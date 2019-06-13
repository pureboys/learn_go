package main

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var (
	esClient *elastic.Client
)

func initES(addr string) (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	esClient = client
	return

	/*
		fmt.Println("conn es success")

		for i := 0; i < 10000; i++ {
			tweet := LogMessage{}
			_, err = client.Index().Index("twitter").Type("tweet").Id(fmt.Sprintf("%d", i)).BodyJson(tweet).Do(context.Background())

			if err != nil {
				panic(err)
				return
			}
		}

		fmt.Println("insert success")
	*/
}

func sendToEs(topic string, data []byte) (err error) {
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)

	_, err = esClient.Index().Index(topic).Type(topic).BodyJson(msg).Do(context.Background())

	if err != nil {
		panic(err)
		return
	}

	return
}
