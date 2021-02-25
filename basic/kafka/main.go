package kafka

import (
	"encoding/base64"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"time"
)

var (
	kafka_addr  = ""
	kafka_topic = ""
)

func main() {
	testConsumer()
}

func testConsumer() {

}

func startConsume(topic, group string) bool {
	// init (custom) config, set mode to ConsumerModePartitions
	cfg := cluster.NewConfig()
	cfg.Consumer.Offsets.Initial = kafka.OffsetOldest
	cfg.Consumer.Fetch.Max = int32(30720000)
	cfg.Consumer.Fetch.Default = int32(30720000)
	cfg.Group.Mode = cluster.ConsumerModePartitions
	cfg.Group.Offsets.Retry.Max = 3
	cfg.Group.Return.Notifications = true
	cfg.Consumer.MaxWaitTime = 10 * time.Second
	cfg.Consumer.MaxProcessingTime = 10 * time.Second

	// init consumer
	topics := []string{kafka_topic}

	log.Printf("start consume topic:%v,group:%v,run topics:%v", topic, group, topics)
	consumer, err := cluster.NewConsumer([]string{kafka_addr}, group, topics, cfg)
	if err != nil {
		logger.Errorf("init consumer is error:%v\ttopic:%v\tgroup:%v", err, topic, group)
		return false
	}
	defer consumer.Close()

	// init redis kafka topic offset to memory
	// if redis not have this key,ignore it and use kafka offset
	rOffsetKey := s.config.String("idc.name") + "_" + s.config.String("platform.id") + "_kafka_offset_" + topic + "_" + group

	if err != nil {
		logger.Errorf("init redis topic offset error:%v\ttopic:%v\tgroup:%v", err, topic, group)
		return false
	}

	//notify rebalance
	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			//TODO 细粒度处理
			logger.Debugf("startConsume Rebalanced topic:%v,group:%v,ntf:%+v,rOffset:%+v,rOffsetKey:%v,err:%v\n", topic, group, ntf, rOffset, rOffsetKey, err)
		}
	}()

	// consume partitions
	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				logger.Errorf("get consumer partitions error:%v\ttopic:%v\tgroup:%v", ok, topic, group)
				return false
			}

			// start a separate goroutine to consume messages
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					//check msg offset which is over the redis offset
					logger.Debugf("get msg from kafka.topic:%s\tgroup:%s\tmsg offset:%v\tmsg partition:%v", topic, group, msg.Offset, msg.Partition)
					checkFlag, errorMsg := s.checkOffset(rOffset, msg.Partition, msg.Offset, rOffsetKey)
					if !checkFlag {
						logger.Errorf("offset is not correct.topic:%s,group:%s,Partition:%v,Offset:%v,errorMsg:%s", topic, group, msg.Partition, msg.Offset, errorMsg)
						consumer.MarkOffset(msg, "")
						continue
					}
					decodeBytes, err := base64.StdEncoding.DecodeString(string(msg.Value))
					if err != nil {
						logger.Errorf("get not a base64 msg is error topic:%v, group:%v,Partition:%v,Offset:%v,err:%v", topic, group, msg.Partition, msg.Offset, err)
					} else {
						//retry,must success and do next
						startTime := time.Now()
						retryTimes := 1
						for {
							runFlag := s.invoker.invokerGateway(decodeBytes, retryTimes, common.GATEWAY_DISPATCH_TYPE_ASYNC_KAFKA)
							if !runFlag {
								logger.Errorf("retry to request service.topic:%v,group:%v,Partition:%v,Offset:%v,retryTimes:%v", topic, group, msg.Partition, msg.Offset, retryTimes)
								if retryTimes >= 5 && s.config.Bool("async.retry") {
									//save to async retry
									flag := s.saveKafkaAsyncRetry(decodeBytes, topic, group, msg.Partition, msg.Offset)
									if flag {
										logger.Errorf("save to async retyr error.topic:%v,group:%v,Partition:%v,Offset:%v,retryTimes:%v", topic, group, msg.Partition, msg.Offset, retryTimes)
										break
									}
								}
								time.Sleep(100 * time.Millisecond)
								retryTimes++
							} else {
								break
							}
						}
						endTime := time.Now()
						usedTime := endTime.Sub(startTime).Nanoseconds() / 1000000
						logger.Debugf("consumer one msg used time, topic:%v,group:%v,mTopic:%+v,Partition:%v,Offset:%v,time:%v,retryTimes:%v", topic, group, msg.Topic, msg.Partition, msg.Offset, usedTime, retryTimes)
					}

					logger.Debugf("write kafka offset to map,topic:%v,group:%v,mTopic:%v,Partition:%v,Offset:%v,", topic, group, msg.Topic, msg.Partition, msg.Offset)
					//commit offset to kafka
					consumer.MarkOffset(msg, "")
				}
			}(part)
		}
	}
	return false
}
