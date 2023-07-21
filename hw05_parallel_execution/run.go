package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errCount struct {
	mu    sync.Mutex
	count int
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	chTasks := make(chan Task)
	chErr := make(chan error)
	errCount := errCount{}

	wgTasks := sync.WaitGroup{}
	wgTasks.Add(1)
	go func() {
		defer wgTasks.Done()
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

	wgWorkers := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wgWorkers.Add(1)
		go func() {
			defer wgWorkers.Done()
			for task := range chTasks {
				errCount.mu.Lock()
				if errCount.count >= m {
					errCount.mu.Unlock()
					return
				}
				errCount.mu.Unlock()
				chErr <- task()
			}
		}()
	}

	wgTasks.Add(1)
	go func() {
		defer wgTasks.Done()
		if m >= 0 {
			for err := range chErr {
				if err != nil {
					errCount.mu.Lock()
					errCount.count++
					errCount.mu.Unlock()
				}
			}
		}
	}()

	wgWorkers.Wait()
	close(chErr)
	wgTasks.Wait()

	if errCount.count >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
