package mq

import (
	"github.com/Shopify/sarama"
	log "github.com/thinkboy/log4go"
)

var (
	producer sarama.AsyncProducer
)

// InitKafka init kafka
func InitKafka(kafkaAddrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err = sarama.NewAsyncProducer(kafkaAddrs, config)
	go handleSuccess(producer)
	go handleError(producer)

	return
}

func handleSuccess(p sarama.AsyncProducer) {
	for {
		pm := <-p.Successes()
		if pm != nil {
			log.Info("producer message success, partition:%d offset:%d key:%v valus:%s", pm.Partition, pm.Offset, pm.Key, pm.Value)
		}
	}
}

func handleError(p sarama.AsyncProducer) {
	for {
		err := <-p.Errors()
		if err != nil {
			log.Error("producer message error, partition:%d offset:%d key:%v valus:%s error(%v)", err.Msg.Partition, err.Msg.Offset, err.Msg.Key, err.Msg.Value, err.Err)
		}
	}
}

// PushKafka push kafka
func PushKafka(topic string, value string) (err error) {
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(value)}
	return
}

// Close close
func Close() {
	if producer == nil {
		return
	}

	err := producer.Close()
	if err != nil {
		log.Error("关闭AsyncProducer时发生错误:%s", err.Error())
	}
}
