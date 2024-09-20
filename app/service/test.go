package service

import (
	"fmt"
	"iv-test/app/model"
	"iv-test/library/kafka"
	"log"

	"github.com/IBM/sarama"
)

// 获取发送时间小于当前时间的消息,并写入kafka
func GetMessage() {
	messageList, _ := model.GetMessageAllList()
	client := kafka.GetKafkaClient()

	for _, v := range messageList {
		fmt.Println("分区数量:", client.Partitions)
		// 发送消息
		partition, offset, err := client.SendMessage(kafka.Topic, "test1", v.Message)
		if err != nil {
			log.Fatalf("Failed to send message: %s", err)
		}
		log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", kafka.Topic, partition, offset)
	}
}

// 接收kafka消息
func ReceiveMessages() {
	client := kafka.GetKafkaClient()
	//接收消息
	go client.ReceiveMessages("test_topic", 1, func(msg *sarama.ConsumerMessage) {
		fmt.Println("Received message: ", string(msg.Key), string(msg.Value))
	})

}
