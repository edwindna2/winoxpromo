package source

import (
	"fmt"
	"log"
	"sync"
)

var (
	once           sync.Once
	singleInstance *Source
)

type Source struct {
	Broker *kafkaSrc
}

func FetchSource() (*Source, error) {
	var err error
	if singleInstance == nil {
		once.Do(
			func() {
				log.Println("Source: Creating single instance now.")
				singleInstance, err = newInstanceSource()
			})
	} else {
		log.Println("Source: Single instance already created.")
	}

	return singleInstance, err
}

func newInstanceSource() (*Source, error) {
	broker, err := newInstanceBrokerSrc()
	if err != nil {
		return nil, fmt.Errorf("Source: error instance - %v", err)
	}

	return &Source{Broker: broker}, nil
}

func (src *Source) Close() {
	if src.Broker != nil {
		src.Broker.close()
	}
}
