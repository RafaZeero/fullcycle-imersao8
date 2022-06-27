package main

import (
	"fmt"
	"log"

	kafka2 "github.com/RafaZeero/fullcycle-fullstack-imersao8/simulador/application/kafka"
	"github.com/RafaZeero/fullcycle-fullstack-imersao8/simulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "readtest", producer)

	// for {
	// 	_ = 1
	// }

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// err := route.LoadPositions()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}
