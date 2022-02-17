package instance

import (
	"github.com/go-magic/work-pool/pool"
	"sync"
)

type instance struct {
	p *pool.WorkerPool
}

var (
	once sync.Once
	in *instance
	MaxRoutine int
)

func GetInstance() *instance {
	once.Do(func() {
		in = &instance{}
		in.p = pool.NewPool(MaxRoutine)
	})
	return in
}

func (i *instance)AddTask(subTask *pool.WorkerTask)  {
	i.p.AddTask(subTask)
}

func (i *instance)AddTasks(subTasks ...*pool.WorkerTask)  {
	for _,subTask := range subTasks {
		i.p.AddTask(subTask)
	}
}
