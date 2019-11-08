package eventbus

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

// Bus exports the interface to communicate through events
type Bus interface {
	Close() error
	Publish(topic string, messages ...*message.Message) error
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
}

// Publisher contains the common structure for publishers
type Publisher interface {
	Close() error
	Publish(topic string, messages ...*message.Message) error
}

// Subscriber contains the common structure for subscribers
type Subscriber interface {
	Close() error
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
}
