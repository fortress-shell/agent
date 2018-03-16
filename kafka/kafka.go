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
	BuildId  string
	Topic    string
	Position int
}

type KafkaStageWriter struct {
	*KafkaWriter
	Stage   string
	Command string
}

type Log struct {
	BuildId  string `json:"build_id"`
	Position int    `json:"position"`
	Content  string `json:"content"`
	Stage    string `json:"stage"`
	Command  string `json:"command"`
}

func (k *KafkaStageWriter) Write(p []byte) (n int, err error) {
	log.Println(string(p))
	logEntry := Log{
		BuildId:  k.BuildId,
		Position: k.Position,
		Content:  string(p),
		Stage:    k.Stage,
		Command:  k.Command,
	}
	b, err := json.Marshal(logEntry)
	if err != nil {
		return 0, err
	}
	msg := &sarama.ProducerMessage{
		Topic: k.Topic,
		Key:   sarama.StringEncoder(k.Id),
		Value: sarama.ByteEncoder(b),
	}
	_, _, err = k.Writer.SendMessage(msg)
	if err != nil {
		return 0, err
	}
	k.Position += 1
	return len(p), nil
}

func NewKafkaWriter(url, topic, id, buildId string) (*KafkaWriter, error) {
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
		Topic:   topic,
	}
	return kafkaWriter, nil
}
