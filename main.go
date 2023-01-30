package main

import (
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/alexperezortuno/go-zmq/commons/structs"
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
		Tester: false,
		Worker: false,
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
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
				worker.Start()
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
