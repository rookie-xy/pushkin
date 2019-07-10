package pushkin

import (
	"testing"
)

func TestProducer(t *testing.T)  {
	config := &ProducerConfig{

	}

	producer, err := NewProducer(config)
	if err != nil {
		t.Log(err.Error())
	}
}