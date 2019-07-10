package pushkin

import (
	"github.com/queueio/pushkin/core/t"
	"github.com/rookie-xy/times33"
)

type Producer struct {
	lc   int64
	conn []*Conn
}

func NewProducer(config *ProducerConfig) (*Producer, error) {
	producer := &Producer{}

	for _, node := range config.Cluster {
		conn, err := NewConn(&ConnConfig{
			Host: node,
		})
		if err != nil {
			// log
			return nil, err
		}
		producer.conn = append(producer.conn, conn)
		producer.lc++
	}

	return producer, nil
}

func (p *Producer) Publish(topics []string, message *t.Message) error {
	for _, topic := range topics {
		c := p.conn[times33.Key(message.Tags)%p.lc]
		// encode message + topic
		c.Write()
	}

	return nil
}
