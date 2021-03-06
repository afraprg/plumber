package rabbitmq

import (
	"context"
	"fmt"

	"github.com/batchcorp/rabbit"
	"github.com/streadway/amqp"

	"github.com/jhump/protoreflect/desc"
	"github.com/pkg/errors"

	"github.com/batchcorp/plumber/cli"
	"github.com/batchcorp/plumber/pb"
	"github.com/batchcorp/plumber/printer"
	"github.com/batchcorp/plumber/reader"
)

// Read is the entry point function for performing read operations in RabbitMQ.
//
// This is where we verify that the provided arguments and flag combination
// makes sense/are valid; this is also where we will perform our initial conn.
func Read(opts *cli.Options) error {
	if err := validateReadOptions(opts); err != nil {
		return errors.Wrap(err, "unable to validate read options")
	}

	var mdErr error
	var md *desc.MessageDescriptor

	if opts.ReadOutputType == "protobuf" {
		md, mdErr = pb.FindMessageDescriptor(opts.ReadProtobufDirs, opts.ReadProtobufRootMessage)
		if mdErr != nil {
			return errors.Wrap(mdErr, "unable to find root message descriptor")
		}
	}

	r, err := New(opts, md)

	if err != nil {
		return errors.Wrap(err, "unable to initialize rabbitmq consumer")
	}

	return r.Read()
}

// Read will attempt to consume one or more messages from the established rabbit
// channel.
func (r *RabbitMQ) Read() error {
	r.log.Info("Listening for message(s) ...")

	errCh := make(chan *rabbit.ConsumeError)
	ctx, cancel := context.WithCancel(context.Background())

	lineNumber := 1

	go r.Consumer.Consume(ctx, errCh, func(msg amqp.Delivery) error {

		data, err := reader.Decode(r.Options, r.MsgDesc, msg.Body)
		if err != nil {
			return err
		}

		str := string(data)

		if r.Options.ReadLineNumbers {
			str = fmt.Sprintf("%d: ", lineNumber) + str
			lineNumber++
		}

		printer.Print(str)

		if !r.Options.ReadFollow {
			cancel()
		}

		return nil
	})

	for {
		select {
		case err := <-errCh:
			return err.Error
		case <-ctx.Done():
			r.log.Debug("Reader exiting")
			return nil
		}
	}

	return nil
}

func validateReadOptions(opts *cli.Options) error {
	if opts.Rabbit.Address == "" {
		return errors.New("--address cannot be empty")
	}

	if opts.Action == "write" && opts.Rabbit.RoutingKey == "" {
		return errors.New("--routing-key cannot be empty with write action")
	}

	if opts.ReadOutputType == "protobuf" {
		if err := cli.ValidateProtobufOptions(
			opts.ReadProtobufDirs,
			opts.ReadProtobufRootMessage,
		); err != nil {
			return fmt.Errorf("unable to validate protobuf option(s): %s", err)
		}
	}

	return nil
}
