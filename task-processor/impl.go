package taskprocessor

import (
	"context"

	"github.com/SimifiniiCTO/asynq"
	"go.uber.org/zap"
)

// EnqueueTask implements IProcessor.
func (tp *TaskProcessor) EnqueueTask(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	if task == nil {
		return nil, ErrTaskNotSet
	}

	tp.logger.Info("enqueueing task", zap.Any("task", task))

	return tp.client.EnqueueContext(ctx, task, opts...)
}

// StartWorker starts the task processor worker
// ```go
//
//			tp, err := NewTaskProcessor(...opts)
//			if err != nil {
//				return err
//			}
//
//	     	// start the worker asynchronously in another go routine
//		 	go tp.StartWorker()
//
// ```
func (tp *TaskProcessor) StartWorker() error {
	// start the worker
	if err := tp.worker.Start(); err != nil {
		return err
	}

	return nil
}

// Close closes the task processor
//
// ```go
//
//			tp, err := NewTaskProcessor(...opts)
//			if err != nil {
//				return err
//			}
//
//			defer tp.Close()
//	     	// start the worker asynchronously in another go routine
//		 	go func(fn TaskProcessorHandler) {
//				tp.StartWorker(fn)
//			}(fn)
//
// ```
func (tp *TaskProcessor) Close() error {
	tp.worker.Stop()

	// close the redis connection
	if err := tp.client.Close(); err != nil {
		return err
	}

	return nil
}

// Validate validates the task processor
func (tp *TaskProcessor) Validate() error {
	if tp.client == nil {
		return ErrClientNotSet
	}

	if tp.worker == nil {
		return ErrWorkerNotSet
	}

	if tp.logger == nil {
		return ErrLoggerNotSet
	}

	if tp.instrumentationClient == nil {
		return ErrInstrumentationClientNotSet
	}

	if tp.concurrencyFactor == nil {
		return ErrConcurrencyFactorNotSet
	}

	if tp.taskHandler == nil {
		return ErrTaskHandlerNotSet
	}

	return nil
}
