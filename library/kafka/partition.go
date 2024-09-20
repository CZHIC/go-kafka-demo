package kafka

import "github.com/IBM/sarama"

// 自定义分区器
type CustomPartitioner struct{}

func (p *CustomPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	// 这里可以根据消息的内容或其他逻辑来决定分区
	// 例如，总是选择第一个分区
	return 0, nil
}

func (p *CustomPartitioner) RequiresConsistency() bool {
	return false
}

// 自定义分区器构造函数
func NewCustomPartitioner(config *sarama.Config) sarama.Partitioner {
	return &CustomPartitioner{}
}
