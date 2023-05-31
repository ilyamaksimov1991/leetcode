package scheduler

import (
	"testing"
	"time"
)

func TestFifo_Run(t *testing.T) {
	sheduller := NewFifo(2)

	sheduller.AddTask(task{
		Id:       "one",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})
	sheduller.AddTask(task{
		Id:       "two",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})
	sheduller.AddTask(task{
		Id:       "three",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})
	sheduller.Run()

	sheduller.AddTask(task{
		Id:       "four",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})

	time.Sleep(time.Second * 5)
}
func TestRR_Run(t *testing.T) {
	sheduller := NewRoundRobin(8, time.Millisecond*50)

	sheduller.AddTask(task{
		Id:       "one",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})
	sheduller.AddTask(task{
		Id:       "two",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})
	sheduller.AddTask(task{
		Id:       "three",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})
	sheduller.Run()

	sheduller.AddTask(task{
		Id:       "four",
		LeedTime: time.Millisecond * 100,
		Priority: 0,
	})

	time.Sleep(time.Second * 15)
}
