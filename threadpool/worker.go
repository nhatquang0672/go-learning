package threadpool

import "fmt"

type Worker struct {
	id          int
	jobQueue    chan interface{}
	closeHandle chan bool
}

func NewWorker(id int, jobQueue chan interface{}, closeHandle chan bool) *Worker {
	return &Worker{id: id, jobQueue: jobQueue, closeHandle: closeHandle}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.jobQueue:
				w.executeJob(job)
			case <-w.closeHandle:
				return
			}
		}
	}()
}

func (w *Worker) executeJob(job interface{}) {
	switch task := job.(type) {
	case Runnable:
		task.Run()
	default:
		fmt.Errorf("Don't know the job type")
	}
}
