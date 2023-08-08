package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageDone(in In, done In, stage Stage) Out {
	temp := make(Bi)
	out := stage(temp)

	go func() {
		defer close(temp)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				temp <- v
			}
		}
	}()

	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	n := len(stages)
	if n == 0 {
		return nil
	}

	pipes := make([]In, 0, n-1)
	for i := 0; i < n-1; i++ {
		pipes = append(pipes, make(In))
	}

	pipes[0] = stageDone(in, done, stages[0])

	for i := 1; i < n-1; i++ {
		pipes[i] = stageDone(pipes[i-1], done, stages[i])
	}

	output := stageDone(pipes[n-2], done, stages[n-1])

	return output
}
