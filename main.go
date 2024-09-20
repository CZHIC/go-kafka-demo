package main

import (
	"fmt"
	_ "iv-test/boot"
	"iv-test/library/kafka"
	_ "iv-test/router"
	"log"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/frame/g"
)

func main() {

	client := kafka.GetKafkaClient()
	fmt.Println("分区数量:", client.Partitions)

	// 发送消息
	partition, offset, err := client.SendMessage("test_topic", "test1", "Hello, Kafka!")
	if err != nil {
		log.Fatalf("Failed to send message: %s", err)
	}
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", "test_topic", partition, offset)

	//接收消息

	client.ReceiveMessages("test_topic", 1, func(msg *sarama.ConsumerMessage) {
		fmt.Println("Received message: ", string(msg.Key), string(msg.Value))
	})

	g.Server().Run()
}
