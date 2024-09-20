package kafka

import (
	"log"

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
		//Partition: partition, //指定分区
	}

	p, offset, err := kc.Producer.SendMessage(msg)
	if err != nil {
		return 0, 0, err
	}

	return p, offset, nil
}

// ReceiveMessages 从Kafka接收消息
func (kc *KafkaClient) ReceiveMessages(topic string, partition int32, handler func(*sarama.ConsumerMessage)) {
	log.Println("进入ReceiveMessages")
	partitionConsumer, err := kc.Consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %s", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		//log.Printf("Received message: key=%s, value=%s, offset=%d\n", string(msg.Key), string(msg.Value), msg.Offset)
		handler(msg)
	}
}
