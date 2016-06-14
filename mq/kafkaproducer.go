package mq

import (
	"strings"

	"github.com/Shopify/sarama"
	log "github.com/thinkboy/log4go"
)

var (
	producer sarama.AsyncProducer
)

func InitKafka(kafkaAddrs []string, quit chan int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err = sarama.NewAsyncProducer(kafkaAddrs, config)
	go handleSuccess(quit)
	go handleError(quit)
	return
}

func handleSuccess(quit chan int) {
	var (
		pm *sarama.ProducerMessage
	)
	for {
		pm = <-producer.Successes()
		if pm != nil {
			log.Info("producer message success, partition:%d offset:%d key:%v valus:%s", pm.Partition, pm.Offset, pm.Key, pm.Value)
		}
		quit <- 1
	}
}

func handleError(quit chan int) {
	var (
		err *sarama.ProducerError
	)
	for {
		err = <-producer.Errors()
		if err != nil {
			log.Error("producer message error, partition:%d offset:%d key:%v valus:%s error(%v)", err.Msg.Partition, err.Msg.Offset, err.Msg.Key, err.Msg.Value, err.Err)
		}
		quit <- 2
	}
}

func pushKafka(topic string, value string) (err error) {
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(value)}
	return
}

//KafkaProducerPush 生产
func KafkaProducerPush(topic string, value string, addr string) {
	var quit = make(chan int)
	kafkaAddrs := strings.Split(addr, ",")
	InitKafka(kafkaAddrs, quit)
	err := pushKafka(topic, value)
	if err != nil {
		log.Info("pushKafka err: ", err)
	}
	handleQuit := <-quit
	log.Info("handleQuit: ", handleQuit)
}
