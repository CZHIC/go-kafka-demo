# 目录介绍
|-- README.MD
|-- app                                          
|   |-- api                     
|       |-- test.go             接口，包含三个接口和一个定时任务
|   |-- service                 
|       |-- test.go             业务逻辑层
|-- config                      配置文件
|-- library
|   |-- kafka                   kafka插件

|-- sql                         存放sql文件
|-- router                       路由

# 实现功能
  1: 定时每分钟查询消息表数据，写入kafka   test.go -> GetMessage()
  2: 导入excel文件，写入消息表          test.go -> UploadFile()   接口 ： /test/upload-file
  3：开启消费端                         test.go -> TestStart()   接口 ： /test/test-start  
  4：关闭消费端                         test.go -> TestStop()    接口 ：  /test/test-stop

# 测试
 1:  根据sql目录里的文件创建数据表 ，
     修改config/config.toml  数据库连接配置
     修改config/config.toml  KAFKA连接配置

 2： 通过/test/upload-file 接口 导入 活动导入.xlsx 文件，将消息写入 t_message 消息表

 3： 定时任务GetMessage() 每分钟执行，根据消息状态，和计划时间 将消息写入 kafka

 4: /test/test-start 接口开启测试消费端

 5：/test/test-stop 关闭消费端

 目前本地测试写入kafka消息 10000+ 条数，开启消费端后，会启动10000+条协程异步处理消息 。可查看图片 =》 性能监控.png


 # kafka 配置
    kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Version = sarama.V2_8_1_0
	kafkaConfig.Producer.Partitioner = sarama.NewHashPartitioner // 按照key hash值分区 

    kafkaConfig.Consumer.Return.Errors = true

# 启动
  go mod tidy
  go run main.go
  访问接口示例 ： 127.0.0.1:8081/test/test-start  （开启消费者）
  系统监控 ： http://localhost:18066/debug/statsview      