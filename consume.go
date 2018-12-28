package easyrabbit

import (
	"github.com/streadway/amqp"
)

type ConsumeCallback func(body []byte)error

func (c *Connection) Consume(queueName string, tag string) (<-chan amqp.Delivery, error) {

	messages, err := c.channel.Consume(
		queueName,  // name
		tag,        // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *Connection) ConsumeWithCallback(queueName string, tag string, cb ConsumeCallback) error {

	messages, err := c.Consume(queueName, tag)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			if err := cb(msg.Body); err == nil {
				msg.Ack(false)
			}
		}
	}()

	return nil
}
