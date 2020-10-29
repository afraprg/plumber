package rabbitmq

import (
	"github.com/batchcorp/rabbit"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/cli"
)

// RabbitMQ holds all attributes required for performing a read/write operations
// in RabbitMQ. This struct should be instantiated via the rabbitmq.Read(..) or
// rabbitmq.Write(..) functions.
type RabbitMQ struct {
	Options  *cli.Options
	Consumer *rabbit.Rabbit
	log      *logrus.Entry
}

func New(opts *cli.Options) (*RabbitMQ, error) {
	mode := rabbit.Consumer
	if opts.Action == "write" {
		mode = rabbit.Producer
	}

	rmq, err := rabbit.New(&rabbit.Options{
		URL:            opts.Rabbit.Address,
		QueueName:      opts.Rabbit.ReadQueue,
		ExchangeName:   opts.Rabbit.Exchange,
		RoutingKey:     opts.Rabbit.RoutingKey,
		QueueExclusive: opts.Rabbit.ReadQueueExclusive,
		AutoAck:        opts.Rabbit.ReadAutoAck,
		ConsumerTag:    opts.Rabbit.ReadConsumerTag,
		AppID:          opts.Rabbit.WriteAppID,
		Mode:           mode,
	})

	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize rabbitmq consumer")
	}

	r := &RabbitMQ{
		Options:  opts,
		Consumer: rmq,
		log:      logrus.WithField("pkg", "rabbitmq.go"),
	}

	return r, nil
}
