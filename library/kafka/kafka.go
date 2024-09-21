package kafka

import (
	"iv-test/library/logger"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

// SendMessage 发送消息到Kafka
func (kc *KafkaClient) SendMessage(topic, key, value string) (int32, int64, error) {
	kc.mu.Lock()
	defer kc.mu.Unlock()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	p, offset, err := kc.Producer.SendMessage(msg)
	if err != nil {
		return 0, 0, err
	}

	return p, offset, nil
}

// ReceiveMessages 从Kafka接收消息
func (kc *KafkaClient) ReceiveMessages(topic string, stopSignal <-chan bool, handler func(*sarama.ConsumerMessage)) {
	log.Println("进入ReceiveMessages")
	var wg sync.WaitGroup
	for _, partition := range kc.Partitions {
		wg.Add(1)
		partitionId := partition.ID
		go func(partitionId int32) {
			defer wg.Done()
			partitionConsumer, err := kc.Consumer.ConsumePartition(topic, partitionId, sarama.OffsetOldest)
			if err != nil {
				logger.Error("Failed to start partition consumer:", err.Error())
				return
			}
			defer partitionConsumer.Close()
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					go handler(msg)
				case err := <-partitionConsumer.Errors():
					logger.Error("kafka获取数据报错:", err.Error())
					return
				case <-stopSignal:
					return
				}
			}

		}(partitionId)

	}
	wg.Wait()
}
