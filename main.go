package main

import (
  "os"
  "fmt"
  "strings"
  "github.com/Shopify/sarama"
)
// import "github.com/hashicorp/memberlist"
// import "github.com/hashicorp/serf"
// import "google.golang.org/grpc"
// import "golang.org/x/crypto/ssh"
// import "
// import "github.com/libvirt/libvirt-go"
// import "github.com/digitalocean/go-qemu"

func main() {
  kafkaUrl := os.Getenv("KAFKA_URL")
  brokerList := strings.Split(kafkaUrl, ",")
  config := sarama.NewConfig()
  config.Producer.RequiredAcks = sarama.WaitForAll
  config.Producer.Retry.Max = 10
  config.Producer.Return.Successes = true
  producer, err := sarama.NewSyncProducer(brokerList, config)
  if err != nil {
    fmt.Println("Failed to start Sarama producer:", err)
  }
  defer producer.Close()
  msg := sarama.ProducerMessage{
    Topic: "dc1",
    Key: sarama.StringEncoder("main"),
    Value: sarama.StringEncoder("top error fuck message"),
  }
  for {
    partition, offset, err := producer.SendMessage(&msg)
    fmt.Println("fuck", partition, offset, err)
  }
}
