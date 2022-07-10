package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Producer interface {
	SendMessage(topic string, body interface{}) error
}

type producer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(kafkaBroker string) Producer {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":     kafkaBroker,
		"broker.address.family": "v4",
		"security.protocol":     "PLAINTEXT",
	})
	if err != nil {
		log.Fatal("Error to connect to kafka", err)
		return nil
	}

	go func() {
		for e := range kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)

				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			case kafka.Error:
				fmt.Printf("Error: %v\n", ev)

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}

		}
	}()

	return &producer{kafkaProducer}
}

func (p producer) SendMessage(topic string, body interface{}) error {
	correlationID := uuid.New().String()

	data := bson.M{
		"body":       body,
		"properties": []int{},
		"headers": bson.M{
			"contentType":   "application/json",
			"correlationId": correlationID,
		},
	}

	source, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error to parser data to byte", err)
		return err
	}

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          source,
		Key:            []byte(correlationID),
		Headers: []kafka.Header{
			{Key: "correlationId", Value: []byte(correlationID)},
			{Key: "applicationName", Value: []byte("ms-miniapp")},
		},
	}, nil)
	if err != nil {
		fmt.Println("Error to sned message to kafka producer", err)
		return err
	}

	p.producer.Flush(-1)
	fmt.Println("Message sent to kafka producer")

	return nil
}
