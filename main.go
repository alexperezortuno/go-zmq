package main

import (
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/alexperezortuno/go-zmq/commons/structs"
	"github.com/alexperezortuno/go-zmq/sink"
	"github.com/alexperezortuno/go-zmq/tester"
	"github.com/alexperezortuno/go-zmq/worker"
	"github.com/pteich/configstruct"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var logger = commons.GetLogger()

	conf := structs.Flags{
		Tester:      false,
		Worker:      false,
		Sink:        false,
		NumOfWorker: 1,
	}

	ctx := make(chan os.Signal)
	signal.Notify(ctx, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ctx
		os.Exit(0)
	}()

	cmd := configstruct.NewCommand(
		"",
		"CLI tool to wun zeromq",
		&conf,
		func(c *configstruct.Command, cfg interface{}) error {
			if conf.Tester {
				tester.Start()
			}

			if conf.Worker {
				worker.Start(cfg.(*structs.Flags))
			}

			if conf.Sink {
				sink.Start()
			}
			return nil
		},
	)

	err := cmd.ParseAndRun(os.Args)
	if err != nil {
		logger.WithField("line", 48).Error(err)
		os.Exit(1)
	}

	os.Exit(0)
}
