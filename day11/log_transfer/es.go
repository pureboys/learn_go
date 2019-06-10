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

func initES() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

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

}
