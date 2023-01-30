package worker

import (
	"encoding/json"
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/alexperezortuno/go-zmq/commons/structs"
	"github.com/pebbe/zmq4"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

var logger = commons.GetLogger()
var nameSpace = "worker"

func Start() {

	context, err := zmq4.NewContext()
	defer func(context *zmq4.Context) {
		err := context.Term()
		if err != nil {
			logger.WithFields(logrus.Fields{"line": 23, "nameSpace": nameSpace}).Error(err)
		}
	}(context)

	if err != nil {
		logger.WithFields(logrus.Fields{"line": 28, "nameSpace": nameSpace}).Error(err)
	}

	receiver, err := context.NewSocket(zmq4.PULL)
	if err != nil {
		logger.WithFields(logrus.Fields{"line": 33, "nameSpace": nameSpace}).Error(err)
	}

	defer func(receiver *zmq4.Socket) {
		err := receiver.Close()
		if err != nil {
			logger.WithFields(logrus.Fields{"line": 39, "nameSpace": nameSpace}).Error(err)
		}
	}(receiver)

	err = receiver.Connect("tcp://localhost:5555")

	if err != nil {
		logger.WithFields(logrus.Fields{"line": 46, "nameSpace": nameSpace}).Error(err)
	}

	logger.WithFields(logrus.Fields{"line": 49, "nameSpace": nameSpace}).Info("Connected to tcp://localhost:5555")

	for {
		var r structs.Request
		msg, err := receiver.Recv(0)
		if err != nil {
			logger.WithFields(logrus.Fields{"line": 53, "nameSpace": nameSpace}).Error(err)
		}

		err = json.Unmarshal([]byte(msg), &r)
		if err != nil {
			logger.WithFields(logrus.Fields{"line": 58, "nameSpace": nameSpace}).Error(err)
		}

		data, err := makeRequest(r.Id)

		if err != nil {
			logger.WithFields(logrus.Fields{"line": 64, "nameSpace": nameSpace}).Error(err)
		}

		logger.WithFields(logrus.Fields{"line": 67, "nameSpace": nameSpace}).Printf("Received %s\n", data.Name)
	}
}

func makeRequest(paramId int) (structs.SWPeople, error) {
	res, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(paramId) + "/")

	if err != nil {
		logger.WithFields(logrus.Fields{"line": 77, "nameSpace": nameSpace}).Error(err)
	}

	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logger.WithFields(logrus.Fields{"line": 81, "nameSpace": nameSpace}).Error(err)
	}

	var responseObject structs.SWPeople
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		logger.WithFields(logrus.Fields{"line": 88, "nameSpace": nameSpace}).Error(err)
	}

	return responseObject, nil
}
