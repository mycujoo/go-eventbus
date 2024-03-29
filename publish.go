package eventbus

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Publish takes some content, encode it to binary and send it to the given topic
func Publish(publisher Publisher, topic string, content interface{}) error {
	encodedContent, err := Encode(content)
	if err != nil {
		return err
	}

	busMessage := message.NewMessage(watermill.NewUUID(), encodedContent)
	err = publisher.Publish(topic, busMessage)
	if err != nil {
		return err
	}

	return nil
}
