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

	for _, task := range tasks {
		if task == nil {
			return ErrErrorsNilTask
		}
	}

	errCount := errCount{}

	chTasks := make(chan Task)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chTasks)
		for _, task := range tasks {
			errCount.mu.Lock()
			if errCount.count >= m {
				errCount.mu.Unlock()
				return
			}
			errCount.mu.Unlock()
			chTasks <- task
		}
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range chTasks {
				err := task()
				if err != nil {
					errCount.mu.Lock()
					errCount.count++
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
