package pool

import "github.com/go-magic/work-pool/task"

/*
worker 协程执行实体
*/
type worker struct {
	taskChan chan *WorkerTask
	exit     chan struct{}
}

/*
NewWorker 创建协程执行实体
*/
func NewWorker(taskChan chan *WorkerTask) *worker {
	w := &worker{
		taskChan: taskChan,
		exit:     make(chan struct{}),
	}
	go w.start()
	return w
}

/*
开启接收并执行任务实体
*/
func (w worker) start() {
	for {
		select {
		case subTask := <-w.taskChan:
			w.do(subTask)
		case <-w.exit:
			return
		}
	}
}

/*
开始执行任务并发送
*/
func (w worker) do(subTask *WorkerTask) {
	subResult := subTask.WorkerFunc(subTask)
	go w.send(subTask, subResult)
}

/*
发送任务
*/
func (w worker) send(subTask *WorkerTask, subResult *task.Result) {
	subTask.ResultChan <- subResult
}

/*
Close 退出协程
*/
func (w worker) Close() {
	w.exit <- struct{}{}
}
