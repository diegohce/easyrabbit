package easyrabbit

import (
	"github.com/streadway/amqp"
	"log"
)

type Connection struct {
	amqpURI     string
	contentType string
	connection  *amqp.Connection
	channel     *amqp.Channel
	notiClose   chan *amqp.Error
}

func New(uri string) (*Connection, error) {

	c := &Connection{
		amqpURI:     uri,
		contentType: "text/plain",
	}

	if err := c.connect(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Connection) Close() {
	if c.connection != nil {
		//close(c.notiClose)
		c.connection.Close()
	}
}


func (c *Connection) connect() error {

	var err error

	if c.connection, err = amqp.Dial(c.amqpURI); err != nil {
		return err
	}

	if c.channel, err = c.connection.Channel(); err != nil {
		c.connection.Close()
		return err
	}

	errorChan := make(chan *amqp.Error)

	c.notiClose = c.connection.NotifyClose(errorChan)

	go func(ch chan *amqp.Error) {

		for amqpErr := range ch {
			log.Printf("Connection closed. %s\n", amqpErr.Reason)
			if err := c.connect(); err != nil {
				log.Printf("Cannot reconnect. %s\n", err.Error())
				//return
			}
		}

	}(c.notiClose)

	return nil
}
