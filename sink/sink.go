package sink

import (
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/pebbe/zmq4"
)

func Start() {
	context, err := zmq4.NewContext()
	var logger = commons.GetLogger()

	defer func(context *zmq4.Context) {
		err := context.Term()
		if err != nil {
			logger.Fatal(err)
		}
	}(context)

	if err != nil {
		logger.Error(err)
	}

	receiver, err := context.NewSocket(zmq4.PULL)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(sender *zmq4.Socket) {
		err := sender.Close()
		if err != nil {
			logger.Error(err)
		}
	}(receiver)

	err = receiver.Bind("tcp://*:5557")
	if err != nil {
		logger.Error(err)
	}

	pub, err := context.NewSocket(zmq4.PUB)
	if err != nil {
		logger.Fatal(err)
	}

	err = pub.Bind("tcp://*:5558")
	if err != nil {
		return
	}

	_, err = pub.Send("Hello", 0)
	if err != nil {
		return
	}

	for {
		msg, err := receiver.RecvMessageBytes(0)
		if err != nil {
			logger.Error(err)
		}

		logger.Debugf("Received message: %s", msg)
	}
}
