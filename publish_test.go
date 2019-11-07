package eventbus

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/stretchr/testify/assert"
)

func TestPublishMessage(t *testing.T) {
	type simpleStruct struct {
		Text    string
		Content []byte
	}

	payload := &simpleStruct{
		Text:    "Hello",
		Content: []byte("Hello"),
	}

	bus := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))

	messages, err := bus.Subscribe(context.Background(), "some.topic")
	assert.Nil(t, err, "Failed to subscribe to topic")

	err = Publish(bus, "some.topic", payload)
	assert.Nil(t, err, "Failed to publish message")

	message := <-messages
	assert.NotNil(t, message, "Failed to receive message")

	target := &simpleStruct{}
	err = Decode(message.Payload, target)
	assert.Nil(t, err, "Failed to decode message")
	assert.Equal(t, payload, target, "Wrong message payload")
}
