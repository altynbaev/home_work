package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrErrorsNilTask = errors.New("a nil task was received")

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

	chTasks := make(chan Task)
	errCount := errCount{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chTasks)
		for _, task := range tasks {
			errCount.mu.Lock()
			if errCount.count >= m {
				errCount.mu.Unlock()
				break
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
				errCount.mu.Lock()
				if errCount.count >= m {
					errCount.mu.Unlock()
					return
				}
				errCount.mu.Unlock()
				if task == nil {
					continue
				}
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
