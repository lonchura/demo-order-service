package wm

import (
	"log"
	"time"
)

type worker struct {
	id int
	queue chan *Message
	handlers chan func()
	signalKill chan int
}

func (w *worker) run() {
	log.Printf("worker(id=%v) start\n", w.id)

	// init
	w.signalKill = make(chan int)

	// loop for chan
	for {
		select {
		// waiting message
		case message := <-w.queue:
			log.Printf("worker(id=%v) recevie message = %v\n", w.id, message)
			time.Sleep(time.Second * 3)
			log.Printf("worker(id=%v) sleep done\n", w.id)
		case handler := <-w.handlers:
			log.Printf("worker(id=%v) handler = %v start\n", w.id, handler)
			handler()
			log.Printf("worker(id=%v) handler = %v finished\n", w.id, handler)
		// kill worker by signal
		case <-w.signalKill:
			log.Printf("worker(id=%v) killed\n", w.id)
			break
		}
	}
}

func (w *worker) kill()  {
	w.signalKill <- 0
}