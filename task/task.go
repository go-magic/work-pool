package task

const (
	StatusErrorCode = 1000
)

type Task struct {
	SubTaskType int         `json:"sub_task_type"`
	SubTaskID   string      `json:"sub_task_id"`
	SubTask     interface{} `json:"sub_task"`
}

func NewTask(taskType int, taskId string, subTask interface{}) *Task {
	return &Task{
		SubTaskType: taskType,
		SubTaskID:   taskId,
		SubTask:     subTask,
	}
}

type Result struct {
	SubStatusCode int         `json:"sub_status_code"`
	SubTaskType   int         `json:"sub_task_type"`
	SubResult     interface{} `json:"sub_result"`
	SubTaskID     string      `json:"sub_task_id"`
	Error         string      `json:"error"`
}

func NewResult(task *Task) *Result {
	return &Result{
		SubTaskType:   task.SubTaskType,
		SubTaskID:     task.SubTaskID,
		SubStatusCode: StatusErrorCode,
	}
}
