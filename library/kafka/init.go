package kafka

import (
	"log"
	"sync"

	"iv-test/library/logger"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/frame/g"
)

// KafkaClient 包含Kafka生产者和消费者的配置信息
type KafkaClient struct {
	Producer   sarama.SyncProducer
	Consumer   sarama.Consumer
	Partitions []*sarama.PartitionMetadata
	mu         sync.Mutex
}

// kafkaConfig 包含Kafka的配置信息
var kafkaConfig = sarama.NewConfig()
var Topic = "test_topic"

// kafkaClient 单例对象，用于保存Kafka连接
var kafkaClient *KafkaClient

var StopSignals chan bool

func init() {
	kafkaClient = &KafkaClient{}
	// 配置生产者
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Version = sarama.V2_8_1_0
	// 设置分区器为默认的分区器
	kafkaConfig.Producer.Partitioner = sarama.NewHashPartitioner // 按照key hash值分区

	kafkaHost := g.Cfg().GetString("KAFKA.Host")
	// 创建生产者
	var err error
	kafkaClient.Producer, err = sarama.NewSyncProducer([]string{kafkaHost}, kafkaConfig)
	if err != nil {
		logger.Error("Failed to start producer: %s", err)
	}

	// 配置消费者
	kafkaConfig.Consumer.Return.Errors = true

	// 创建消费者
	kafkaClient.Consumer, err = sarama.NewConsumer([]string{kafkaHost}, kafkaConfig)
	if err != nil {
		logger.Error("Failed to start consumer: %s", err)
	}
	//获取分区数量
	admin, err := sarama.NewClusterAdmin([]string{kafkaHost}, kafkaConfig)
	if err != nil {
		logger.Error("failed to create cluster admin: %w", err)
	}
	defer admin.Close()

	describeTopicsResponse, err := admin.DescribeTopics([]string{Topic})
	if err != nil {
		logger.Error("failed to get partitions for topic %s: %w", Topic, err)
	}
	if len(describeTopicsResponse) == 0 {
		logger.Error("failed to get describeTopicsResponse for topic %s: %w", Topic, err)
	}
	partitions := describeTopicsResponse[0].Partitions
	kafkaClient.Partitions = partitions
	log.Println("Kafka connection initialized.")

	StopSignals = make(chan bool, 1)

}

// GetKafkaClient 获取Kafka客户端单例
func GetKafkaClient() *KafkaClient {
	return kafkaClient
}
