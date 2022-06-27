package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	route2 "github.com/RafaZeero/fullcycle-fullstack-imersao8/simulador/application/route"
	"github.com/RafaZeero/fullcycle-fullstack-imersao8/simulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println("Error exporting")
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}