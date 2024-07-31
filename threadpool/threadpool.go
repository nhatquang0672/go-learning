package threadpool

import "fmt"

var (
	ErrQueueFull = fmt.Errorf("queue is full, not able add the task")
)

type ThreadPool struct {
	noOfWorkers int
	queueSize   int64
	jobQueue    chan interface{}
	closeHandle chan bool
}

func NewThreadPool(noOfWorkers int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{noOfWorkers: noOfWorkers, queueSize: queueSize}
	threadPool.jobQueue = make(chan interface{}, queueSize)
	threadPool.closeHandle = make(chan bool)
	threadPool.initWorkers()
	return threadPool
}

func (tp *ThreadPool) initWorkers() {
	for i := 0; i < tp.noOfWorkers; i++ {
		worker := NewWorker(i, tp.jobQueue, tp.closeHandle)
		worker.Start()
	}
}

func (tp *ThreadPool) Close() {
	close(tp.jobQueue)
	close(tp.closeHandle)
}

func (tp *ThreadPool) submitTask(task interface{}) error {
	if len(tp.jobQueue) == int(tp.queueSize) {
		return ErrQueueFull
	}
	tp.jobQueue <- task
	return nil
}

func (tp *ThreadPool) Execute(task Runnable) error {
	return tp.submitTask(task)
}
