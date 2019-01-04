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
