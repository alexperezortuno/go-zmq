package worker

import (
	"encoding/json"
	"fmt"
	"github.com/alexperezortuno/go-zmq/commons"
	"github.com/alexperezortuno/go-zmq/commons/structs"
	zmq "github.com/pebbe/zmq4"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var logger = commons.GetLogger()

func workers(wg *sync.WaitGroup, ctx *zmq.Context, worker int) {
	defer wg.Done()
	//  Socket to receive messages on
	receiver, err := ctx.NewSocket(zmq.PULL)

	if err != nil {
		logger.Error(err)
	}

	err = receiver.Connect("tcp://localhost:5555")

	if err != nil {
		logger.Error(err)
	}

	//  Socket to send messages to task sink
	sender, err := ctx.NewSocket(zmq.PUSH)
	if err != nil {
		logger.Error(err)
	}

	err = sender.Connect("tcp://localhost:5557")
	if err != nil {
		logger.Error(err)
	}

	// Socket to send messages to task sink
	subscriber, err := ctx.NewSocket(zmq.SUB)
	defer func(subscriber *zmq.Socket) {
		err := subscriber.Close()
		if err != nil {

		}
	}(subscriber)

	if err != nil {
		logger.Error(err)
	}

	err = subscriber.Connect("tcp://localhost:5558")
	if err != nil {
		logger.Error(err)
	}

	err = subscriber.SetSubscribe("")
	if err != nil {
		logger.Error(err)
	}

	for {
		var r structs.Request
		msg, err := receiver.Recv(0)
		if err != nil {
			logger.Error(err)
		}

		err = json.Unmarshal([]byte(msg), &r)
		if err != nil {
			logger.Error(err)
		}

		data, err := makeRequest(r.Id)

		if err != nil {
			logger.Error(err)
		}

		logger.Debug("Received ", data.Name)

		taskCompleted := fmt.Sprintf("Worker %d completed task", worker)

		_, err = sender.Send(taskCompleted, 0)
		if err != nil {
			return
		}
	}
}

func Start(conf *structs.Flags) {
	context, err := zmq.NewContext()
	var wg sync.WaitGroup
	wg.Add(conf.NumOfWorker)

	if err != nil {
		logger.Error(err)
	}

	for i := 0; i < conf.NumOfWorker; i++ {
		logger.Debug("Starting worker ", i)
		go workers(&wg, context, i)
	}
	wg.Wait()
}

func makeRequest(paramId int) (structs.SWPeople, error) {
	res, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(paramId) + "/")

	if err != nil {
		logger.Error(err)
	}

	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logger.Error(err)
	}

	var responseObject structs.SWPeople
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		logger.Error(err)
	}

	return responseObject, nil
}
