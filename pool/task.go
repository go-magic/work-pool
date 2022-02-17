package pool

import "github.com/go-magic/work-pool/task"

type WorkerFunc func(*WorkerTask) *task.Result

type WorkerTask struct {
	task.Task
	WorkerFunc WorkerFunc        `json:"worker_func"`
	ResultChan chan *task.Result `json:"result_chan"`
}

func NewWorkerTask(task task.Task, workerFunc WorkerFunc, resultChan chan *task.Result) *WorkerTask {
	return &WorkerTask{
		Task:       task,
		WorkerFunc: workerFunc,
		ResultChan: resultChan,
	}
}
