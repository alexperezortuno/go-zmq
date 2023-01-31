package tester

import (
	"encoding/json"
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/alexperezortuno/go-zmq/commons/structs"
	"github.com/pebbe/zmq4"
	"math/rand"
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

	sender, err := context.NewSocket(zmq4.PUSH)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(sender *zmq4.Socket) {
		err := sender.Close()
		if err != nil {
			logger.Error(err)
		}
	}(sender)

	err = sender.Bind("tcp://*:5555")
	if err != nil {
		return
	}

	for i := 0; i < 20; i++ {
		m := structs.Request{Id: rand.Intn(100)}

		b, err := json.Marshal(m)

		if err != nil {
			logger.Error(err)
		}

		_, err = sender.SendMessage(b)
		if err != nil {
			return
		}

		logger.Infof("Send message %d\n", i)
	}
}
