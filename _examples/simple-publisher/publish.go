// This example declares a durable Exchange, an ephemeral (auto-delete) Queue,
// binds the Queue to the Exchange with a binding key, and consumes every
// message published to that Exchange with that routing key.
//
package main

import (
	"github.com/diegohce/easyrabbit"
)

func main() {

	erc, _ := easyrabbit.New("amqp://guest:guest@localhost/demo")
	defer erc.Close()

	messageBody := "Hello, World!"

	erc.Publish("myExchangeName", "easyrabbit pub demo", []byte(message))
}
