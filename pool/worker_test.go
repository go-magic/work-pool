package pool

import (
	"fmt"
	"github.com/go-magic/work-pool/task"
	"net/http"
	"sync"
	"testing"
)

func Request(workTask *WorkerTask) *task.Result {
	result := task.NewResult(&workTask.Task)
	url, ok := workTask.SubTask.(string)
	if !ok {
		return result
	}
	res, err := http.Get(url)
	if err != nil {
		result.Error = err.Error()
		return result
	}
	result.SubStatusCode = res.StatusCode
	return result
}

type HttpRequest struct {
	resultChan chan *task.Result
	wg         sync.WaitGroup
	endChan    chan struct{}
}

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{
		resultChan: make(chan *task.Result),
		endChan:    make(chan struct{}),
	}
}

type Add struct {
	A int
	B int
}

func AddFunc(workTask *WorkerTask) *task.Result {
	result := task.NewResult(&workTask.Task)
	addData, ok := workTask.SubTask.(*Add)
	if !ok {
		return result
	}
	result.SubResult = addData.A + addData.B
	result.SubStatusCode = 200
	return result
}

func (h *HttpRequest) createTasks() []*WorkerTask {
	tasks := make([]*WorkerTask, 0)
	webTask := NewWorkerTask(
		*task.NewTask(1, "", "https://www.qq.com"),
		Request, h.resultChan)
	tasks = append(tasks, webTask)
	webTask = NewWorkerTask(
		*task.NewTask(2, "", &Add{1, 2}),
		AddFunc, h.resultChan)
	tasks = append(tasks, webTask)
	return tasks

}

func (h *HttpRequest) addWaitGroup() {
	h.wg.Add(1)
}

func (h *HttpRequest) waitGroup() {
	h.wg.Wait()
	h.endChan <- struct{}{}
}

func (h *HttpRequest) wait() {
	for {
		select {
		case result := <-h.resultChan:
			fmt.Println(result)
			h.wg.Done()
		case <-h.endChan:
			return
		}
	}
}

func TestNewPool(t *testing.T) {
	request := NewHttpRequest()
	tasks := request.createTasks()
	p := NewPool(10)
	for _, subTask := range tasks {
		request.addWaitGroup()
		p.AddTask(subTask)
	}
	go request.waitGroup()
	request.wait()
}
