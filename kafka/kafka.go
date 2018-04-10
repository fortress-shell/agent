package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"strings"
)

type KafkaWriter struct {
	Writer   sarama.SyncProducer
	Id       string
	BuildId  int
	UserId   int
	Topic    string
	Position int
}

type KafkaStageWriter struct {
	*KafkaWriter
}

type Log struct {
	BuildId  int    `json:"build_id"`
	UserId   int    `json:"user_id"`
	Position int    `json:"position"`
	Content  string `json:"content"`
}

func (k *KafkaStageWriter) Write(p []byte) (n int, err error) {
	go log.Println(string(p))
	logEntry := Log{
		BuildId:  k.BuildId,
		UserId:   k.UserId,
		Position: k.Position,
		Content:  string(p),
	}
	b, err := json.Marshal(logEntry)
	if err != nil {
		return 0, err
	}
	_, _, err = k.Writer.SendMessage(&sarama.ProducerMessage{
		Topic: k.Topic,
		Key:   sarama.StringEncoder(k.Id),
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		return 0, err
	}
	k.Position += 1
	return len(p), nil
}

func NewKafkaWriter(url, topic, id string, buildId, userId int) (*KafkaWriter,
	error) {
	brokerList := strings.Split(url, ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}
	kafkaWriter := &KafkaWriter{
		Writer:  producer,
		Id:      id,
		BuildId: buildId,
		UserId:  userId,
		Topic:   topic,
	}
	return kafkaWriter, nil
}
