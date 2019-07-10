package pushkin

type Producer struct {
}

func NewProducer(config *ProducerConfig) (*Producer, error) {
	return &Producer{

	}, nil
}

func (p *Producer) Publish() error {
	return nil
}
