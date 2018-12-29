package easyrabbit

import (
	"github.com/streadway/amqp"
	"log"
)

// Connection represents a connection to an AMQP server.
type Connection struct {
	amqpURI     string
	contentType string
	connection  *amqp.Connection
	channel     *amqp.Channel
	notiClose   chan *amqp.Error
}

// New constructs a new connection with the given AMQP Uri.
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

// Close closes an established connection.
func (c *Connection) Close() {
	if c.connection != nil {
		//close(c.notiClose)
		c.connection.Close()
	}
}

// SetContentType sets AMQP content type attribute.
// Default: "text/plain"
func (c *Connection) SetContentType(contentType string) {
	c.contentType = contentType
}

// ContentType returns the content type this connection will use to send messages to the AMQP server.
func (c *Connection) ContentType() string {
	return c.contentType
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
