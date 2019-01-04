// Copyright (c) 2019, Diego Cena.
// Use of this source code is governed by an AGPL-3.0
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/diegohce/easyrabbit

/*
Package easyrabbit is a simplified RabbitMQ client encapsulation
of https://github.com/streadway/amqp package.

Use cases are the simple and most common tasks of publish / consume 
messages from / to RabbitMQ server.

Understand the AMQP 0.9.1 messaging model by reviewing these links first.
As well as the underlaying amqp package.

  Resources

  https://github.com/streadway/amqp
  http://www.rabbitmq.com/tutorials/amqp-concepts.html
  http://www.rabbitmq.com/getstarted.html
  http://www.rabbitmq.com/amqp-0-9-1-reference.html

Use Case

It's important as a client to an AMQP topology to ensure the state of the
broker matches your expectations.  For both publish and consume use cases,
make sure you declare the queues, exchanges and bindings you expect to exist
prior to calling Connection.Publish or Connection.Consume 
or Connection.ConsumeWithCallback.

  // Connections start with easyrabbit.New() typically from a command line argument
  // or environment variable.
  connection, err := easyrabbit.New(os.Getenv("AMQP_URL"))

  // To cleanly shutdown by flushing kernel buffers, make sure to close and
  // wait for the response.
  defer connection.Close()

SSL/TLS - Secure connections

Supported by the underlying package, but not fully implemented on easyrabbit (yet).
*/
package easyrabbit
