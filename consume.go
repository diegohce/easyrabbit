package easyrabbit

import (
	"github.com/streadway/amqp"
)

// ConsumeCallback type for callback function on ConsumeWithCallback.
// Returned error:
// if error == nil, message will be AKed to the broker.
// else, message will be UNAKed.
type ConsumeCallback func(body []byte) error

// Consume starts consumming from queueName identified as tag.
// Returns an amqp.Delivery type channel to consume messages from.
func (c *Connection) Consume(queueName string, tag string) (<-chan amqp.Delivery, error) {

	messages, err := c.channel.Consume(
		queueName, // name
		tag,       // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// ConsumeWithCallback starts consumming from queueName identified as tag.
// Every message content is passed to cb.
//
// Send 'true' to the returned channel to stop the consumer goroutine.
func (c *Connection) ConsumeWithCallback(queueName, tag string, cb ConsumeCallback) (chan bool, error) {

	stopConsumer := make(chan bool)

	messages, err := c.Consume(queueName, tag)
	if err != nil {
		return nil, err
	}

	/*go func() {
		for msg := range messages {
			if err := cb(msg.Body); err == nil {
				msg.Ack(false)
			}
		}
	}()*/

	go func() {
		for {
			select {
			case stop := <-stopConsumer:
				if stop {
					return
				}
			case msg := <-messages:
				{
					if err := cb(msg.Body); err == nil {
						msg.Ack(false)
					}
				}
			}
		}
	}()

	return stopConsumer, nil
}
