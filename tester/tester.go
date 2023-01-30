package tester

import (
	"encoding/json"
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/alexperezortuno/go-zmq/commons/structs"
	"github.com/pebbe/zmq4"
	"github.com/sirupsen/logrus"
	"math/rand"
)

func Start() {
	context, err := zmq4.NewContext()
	var logger = commons.GetLogger()
	var nameSpace = "tester"
	defer func(context *zmq4.Context) {
		err := context.Term()
		if err != nil {
			logger.WithFields(logrus.Fields{"line": 19, "nameSpace": nameSpace}).Fatal(err)
		}
	}(context)

	if err != nil {
		logger.WithFields(logrus.Fields{"line": 24, "nameSpace": nameSpace}).Error(err)
	}

	sender, err := context.NewSocket(zmq4.PUSH)
	if err != nil {
		logger.WithFields(logrus.Fields{"line": 29, "nameSpace": nameSpace}).Fatal(err)
	}

	defer func(sender *zmq4.Socket) {
		err := sender.Close()
		if err != nil {
			logger.WithFields(logrus.Fields{"line": 35, "nameSpace": nameSpace}).Error(err)
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
			logger.WithFields(logrus.Fields{"line": 50, "nameSpace": nameSpace}).Error(err)
		}

		_, err = sender.SendMessage(b)
		if err != nil {
			return
		}

		logger.WithFields(logrus.Fields{"line": 58, "nameSpace": nameSpace}).Infof("Send message %d\n", i)
	}
}
