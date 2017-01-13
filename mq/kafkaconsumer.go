package mq

import (
	"fmt"
	llog "log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	// log "github.com/thinkboy/log4go"
	"github.com/wvanbergen/kafka/consumergroup"
)

// KafkaConsumer kafka consumer
type KafkaConsumer struct {
	Topic             string
	GoupName          string
	ZKRoot            string
	ZKAddrs           []string
	ProcessingTimeout time.Duration
	CommitInterval    time.Duration
}

// InitKafkaConsumer run
func (kc *KafkaConsumer) InitKafkaConsumer(prehookF func(), handleF func(string), posthookF func()) error {
	// log.Info("start topic:%s consumer", kc.Topic)
	// log.Info("consumer group name:%s", kc.GoupName)
	beego.Info(fmt.Sprintf("start topic:%s consumer", kc.Topic))
	beego.Info(fmt.Sprintf("consumer group name:%s", kc.GoupName))
	sarama.Logger = llog.New(os.Stdout, "[Sarama] ", llog.LstdFlags)
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = kc.ProcessingTimeout
	config.Offsets.CommitInterval = kc.CommitInterval
	config.Zookeeper.Chroot = kc.ZKRoot
	kafkaTopics := []string{kc.Topic}

	cg, err := consumergroup.JoinConsumerGroup(kc.GoupName, kafkaTopics, kc.ZKAddrs, config)
	if err != nil {
		return err
	}

	//prehook func
	prehookF()

	go func() {
		for err := range cg.Errors() {
			// log.Error("consumer error(%v)", err)
			beego.Error(fmt.Sprintf("consumer error(%v)", err))
		}
	}()

	go func() {
		for msg := range cg.Messages() {
			// log.Info("deal with topic:%s, partitionId:%d, Offset:%d, Key:%s msg:%s ", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			// log.Info("handle msg: % s ", msg.Value)
			beego.Info(fmt.Sprintf("deal with topic:%s, partitionId:%d, Offset:%d, Key:%s msg:%s ", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value))
			beego.Info(fmt.Sprintf("handle msg: % s ", msg.Value))

			//handle 消费消息
			handleF(string(msg.Value))

			cg.CommitUpto(msg)
		}
	}()

	//posthook func
	posthookF()

	return nil
}

/* e.g
func prehookF() {
	fmt.Println("prehook")
}

func handleF(msg string) {
	fmt.Println(msg)
}

func posthookF() {
	fmt.Println("posthook")
}

func main() {
	topic := "cron"
	groupname := "go"
	zkroot := ""
	zkaddrs := "172.16.9.4:2181,172.16.9.4:2181"
	pt := 10
	ci := 10

	var kafkaConsumer core.KafkaConsumer
	kafkaConsumer.Topic = topic
	kafkaConsumer.GoupName = groupname
	kafkaConsumer.ZKRoot = zkroot
	kafkaConsumer.ZKAddrs = strings.Split(zkaddrs, ",")
	kafkaConsumer.ProcessingTimeout = time.Duration(pt) * time.Second
	kafkaConsumer.CommitInterval = time.Duration(ci) * time.Second
	kafkaConsumer.InitKafkaConsumer(prehookF, handleF, posthookF)

	for {
	}
}
*/
