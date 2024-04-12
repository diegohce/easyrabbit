package easyrabbit

import (
	//"github.com/streadway/amqp"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Publish send body to the AMQP server, publishin into exchange using routingKey.
func (c *Connection) Publish(exchange, routingKey string, body []byte) error {

	err := c.channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     c.contentType,
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient,
			Priority:        0,
		},
	)
	return err
}
