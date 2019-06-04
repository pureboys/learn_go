package main

import (
	"context"
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("localhost:9200"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("conn es success")

	tweet := Tweet{User: "olivere", Message: "Take five"}
	_, err = client.Index().Index("twitter").Type("tweet").Id("1").BodyJson(tweet).Do(context.Background())

	if err != nil {
		panic(err)
		return
	}

	fmt.Println("insert success")

}
