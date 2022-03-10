package kafka_test

import (
	"context"

	"yangcong/modules/kafka"
)

func ExampleWriter() {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "Topic-1",
	}

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
	)

	w.Close()
}
