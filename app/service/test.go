package service

import (
	"encoding/json"
	"fmt"
	"iv-test/app/model"
	"iv-test/library/kafka"
	"iv-test/library/logger"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

// 获取发送时间小于当前时间的消息，并且状态等于未发送,并写入kafka
func GetMessage() {
	total, _ := model.GetMessageNumber()
	if total < 1 {
		return
	}
	logger.Println("total:", total)
	messageList, _ := model.GetMessageAllList()
	logger.Println("消息列表：", messageList)
	client := kafka.GetKafkaClient()
	var wg sync.WaitGroup
	for _, v := range messageList {
		// 发送消息
		wg.Add(1)
		go func() {
			defer wg.Done()
			jsonStr, err := json.Marshal(v)
			if err != nil {
				logger.Error("json错误:", err.Error())
				return
			}
			partition, offset, err := client.SendMessage(kafka.Topic, "test1", string(jsonStr))
			if err != nil {
				log.Fatalf("Failed to send message: %s", err)
			}
			log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", kafka.Topic, partition, offset)
			//更新消息状态
			model.UpdateMessageIfSend(v.Id, 1)
		}()

	}
	wg.Wait()
}

// 测试消费者
func TestStart() {
	client := kafka.GetKafkaClient()

	go client.ReceiveMessages(kafka.Topic, kafka.StopSignals, func(msg *sarama.ConsumerMessage) {
		messageIfo := &model.TMessage{}

		json.Unmarshal([]byte(msg.Value), messageIfo)
		if messageIfo.Id == 0 {
			return
		}
		fmt.Println("接收到消息：: ", string(msg.Key), string(messageIfo.Message))
		//todo 消费操作
		logger.Println("当前消息ID :", messageIfo.Id, "  开始睡眠20s")
		time.Sleep(time.Second * 20)
		logger.Println("当前消息ID :", messageIfo.Id, "  睡眠结束，关闭协程")
		//更新消息状态
		//model.UpdateMessageIfSend(messageIfo.Id, 2)

	})
	return
}

// 发送关闭信号，关闭消费者协程
func TestStop() {
	client := kafka.GetKafkaClient()
	for _, partition := range client.Partitions {
		logger.Println("发送关闭信号:", partition.ID)
		kafka.StopSignals <- true
	}

}
