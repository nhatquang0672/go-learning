package threadpool

import (
	"fmt"
	"testing"
	"time"
)

const (
	NumberOfWorkers = 20
	QueueSize       = int64(1000)
)

var (
	threadpool *ThreadPool
)

func TestNewThreadPool(t *testing.T) {
	threadpool = NewThreadPool(NumberOfWorkers, QueueSize)
}

func TestThreadPool_Execute(t *testing.T) {
	for i := 0; i < 100; i++ {
		data := &TestData{Val: "pristine", Id: i}
		task := &TestTask{TestData: data}
		threadpool.Execute(task)
	}
	time.Sleep(10 * time.Second)
}

func TestThreadPool_Close(t *testing.T) {
	threadpool.Close()
}

func TestQueueFullError(t *testing.T) {
	threadpool := NewThreadPool(0, 1)

	data := &TestData{Val: "pristine"}
	task := &TestTask{TestData: data}

	err := threadpool.Execute(task)
	if err != nil {
		t.Fail()
	}

	err = threadpool.Execute(task)
	if err == nil || err != ErrQueueFull {
		t.Fail()
	}

	threadpool.Close()
}

type TestTask struct {
	TestData *TestData
}

type TestData struct {
	Val string
	Id  int
}

func (t *TestTask) Run() {
	fmt.Printf("Running the task %v \n", t.TestData.Id)
	time.Sleep(1 * time.Second)
	t.TestData.Val = "changed"
}

type TestLongTask struct{}

func (t TestLongTask) Run() {
	time.Sleep(5 * time.Second)
}
