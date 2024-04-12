package easyrabbit_test

import (
	"github.com/diegohce/easyrabbit"
)

func ExampleConnection_Publish() {
	erc, err := easyrabbit.New("amqp://guest:guest@localhost/demo")
	if err != nil {
		// do something with err
	}
	defer erc.Close()

	messageBody := "Hello, World!"

	erc.Publish("myExchangeName", "easyrabbit pub demo", []byte(messageBody))
}

func ExampleConnection_SetContentType() {
	erc, err := easyrabbit.New("amqp://guest:guest@localhost/demo")
	if err != nil {
		// do something with err
	}
	defer erc.Close()

	erc.SetContentType("application/json")
}

func ExampleConnection_ConsumeWithCallback() {
	erc, err := easyrabbit.New("amqp://guest:guest@localhost/demo")
	if err != nil {
		// do something with err
	}
	defer erc.Close()

	// Assuming messageCallback is defined as:
	//func messageCallback(body []byte) error {
	//	fmt.Println(string(body))
	//}
	stopConsumer, err := erc.ConsumeWithCallback("myQueueName", "easyrabbit examples", messageCallback)
	if err != nil {
		// do something with err
	}

	// To stop consumer
	stopConsumer <- true
}
