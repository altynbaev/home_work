package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrErrorsNilTask = errors.New("a nil task was passed")

type Task func() error

type errCount struct {
	mu    sync.Mutex
	count int
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	for _, task := range tasks {
		if task == nil {
			return ErrErrorsNilTask
		}
	}

	errCount := errCount{}

	chTasks := make(chan Task, len(tasks))
	go func() {
		defer close(chTasks)
		for _, task := range tasks {
			chTasks <- task
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range chTasks {
				err := task()
				if err != nil {
					errCount.mu.Lock()
					errCount.count++
					if errCount.count >= m {
						errCount.mu.Unlock()
						return
					}
					errCount.mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	if errCount.count >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
