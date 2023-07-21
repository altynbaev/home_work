package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	chTasks := make(chan Task)
	chErr := make(chan error)
	chStop := make(chan struct{}, 1)
	var errCount int32
	k := int32(m)

	wgTasks := sync.WaitGroup{}
	wgTasks.Add(1)
	go func() {
		defer wgTasks.Done()
		defer close(chTasks)
		for _, task := range tasks {
			select {
			case <-chStop:
				return
			default:
				chTasks <- task
			}
		}
	}()

	wgWorkers := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wgWorkers.Add(1)
		go func() {
			defer wgWorkers.Done()
			for task := range chTasks {
				if errCount >= k {
					return
				}
				chErr <- task()
				// err := task()
				// if err != nil {
				// 	chErr <- err
				// }
			}
		}()
	}

	wgTasks.Add(1)
	go func() {
		defer wgTasks.Done()
		if m >= 0 {
			for err := range chErr {
				if err != nil {
					atomic.AddInt32(&errCount, 1)
					if errCount == k {
						chStop <- struct{}{}
						close(chStop)
					}
				}
			}
		}
	}()

	go func() {
		wgWorkers.Wait()
		close(chErr)
	}()

	wgTasks.Wait()

	if errCount >= k {
		return ErrErrorsLimitExceeded
	}

	return nil
}
