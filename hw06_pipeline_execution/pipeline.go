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

	output := stageDone(in, done, stages[0])
	for i := 1; i < n; i++ {
		output = stageDone(output, done, stages[i])
	}

	return output
}
