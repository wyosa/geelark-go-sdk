package browserautomation

const (
	// TaskStatusWaiting means the task is waiting for its scheduled execution.
	TaskStatusWaiting = 1
	// TaskStatusRunning means the task is executing.
	TaskStatusRunning = 2
	// TaskStatusCompleted means the task completed successfully.
	TaskStatusCompleted = 3
	// TaskStatusFailed means the task failed; inspect ResultCode and ResultDesc.
	TaskStatusFailed = 4
	// TaskStatusCancelled means the task was cancelled.
	TaskStatusCancelled = 7
)
