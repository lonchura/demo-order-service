package wm

import (
	"log"
	"sync"
)

const workerNums = 10
const queueBacklog = 1024

var wg = sync.WaitGroup{}
var WorkerManager *workerManager

type Message struct {
	data interface{}
}

type workerManager struct {
	queue chan *Message
	handlers chan func()
	workers []*worker
}

func NewWorkerManager() *workerManager {
	WorkerManager = &workerManager{}
	return WorkerManager
}

func CreateMessage(data interface{}) *Message {
	return &Message{data: data}
}

func (wm *workerManager) SendMessage(message *Message)  {
	wm.queue <- message
}

func (wm *workerManager) Async(handler func())  {
	wm.handlers <- handler
}

func (wm *workerManager) Start()  {
	// init
	wm.queue = make(chan *Message, queueBacklog)
	wm.handlers = make(chan func())
	wm.workers = make([]*worker, workerNums)

	// start worker
	for i:=0; i<workerNums; i++ {
		// add to wg
		wg.Add(1)
		// add to list
		wm.workers[i] = &worker{
			id: i,
			queue: wm.queue,
			handlers: wm.handlers,
		}
		// run worker
		go wm.workers[i].run()
	}

	log.Printf("WorkerManager start\n")
}

func (wm *workerManager) Stop()  {
	// send stop signal
	for i:=0; i<workerNums; i++ {
		if wm.workers[i] != nil {
			wm.workers[i].kill()
		}
	}

	// wait
	wg.Wait()

	log.Printf("WorkerManager stop\n")
}