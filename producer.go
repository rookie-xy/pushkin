package pushkin

import "github.com/queueio/pushkin/core/t"

type Producer struct {
}

func NewProducer(config *ProducerConfig) (*Producer, error) {
	return &Producer{

	}, nil
}

func (p *Producer) Publish(topics []string, message *t.Message) error {
	return nil
}
