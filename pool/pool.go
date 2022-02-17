package pool

import "github.com/go-magic/work-pool/task"

type workerPool struct {
	workerTaskChan chan *WorkerTask
	workers        chan *worker
}

func NewPool(routine int) *workerPool {
	p := &workerPool{}
	p.initWorkers(routine)
	return p
}

/*
初始化workers
*/
func (p *workerPool) initWorkers(routine int) {
	p.workers = make(chan *worker, routine)
	p.initChan(routine)
	for i := 0; i < routine; i++ {
		p.workers <- NewWorker(p.workerTaskChan)
	}
}

/*
初始化worker chan
*/
func (p *workerPool) initChan(routine int) {
	p.workerTaskChan = make(chan *WorkerTask, routine)
}

/*
AddTask 增加缓冲
*/
func (p *workerPool) AddTask(task *WorkerTask) {
	go func() {
		p.workerTaskChan <- task
	}()
}

/*
WaitResult 等待任务返回,由提供任务的对象提供返回channel
*/
func (p *workerPool) WaitResult(task *WorkerTask) <-chan *task.Result {
	return task.ResultChan
}
