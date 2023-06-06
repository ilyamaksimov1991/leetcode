package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	Id       string
	LeedTime time.Duration
	Priority int
}

type Fifo struct {
	workersCount int
	tasks        chan task
}

func NewFifo(workersCount int) *Fifo {
	return &Fifo{
		workersCount: workersCount,
		tasks:        make(chan task, 100),
	}
}

func (f *Fifo) AddTask(task task) {
	f.tasks <- task
}

func (f *Fifo) Run() {
	go func() {
		wg := sync.WaitGroup{}
		for i := 1; i < f.workersCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for task := range f.tasks {
					fmt.Printf("%d %s start ", i, task.Id)
					time.Sleep(task.LeedTime)
					fmt.Printf("%s stop\n", task.Id)
				}
			}(i)
		}
		wg.Wait()
	}()
}

type RoundRobin struct {
	workersCount int
	timeWork     time.Duration
	tasks        chan task
}

func NewRoundRobin(workersCount int, timeWork time.Duration) *RoundRobin {
	return &RoundRobin{
		workersCount: workersCount,
		timeWork:     timeWork,
		tasks:        make(chan task, 100),
	}
}

func (f *RoundRobin) AddTask(task task) {
	f.tasks <- task
}

func (f *RoundRobin) Run() {
	go func() {
		wg := sync.WaitGroup{}
		for i := 1; i < f.workersCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for task := range f.tasks {
					fmt.Printf("%d %s start \n", i, task.Id)
					time.Sleep(task.LeedTime)
					fmt.Printf("%s stop \n", task.Id)
					if task.LeedTime > f.timeWork {
						leedTime := task.LeedTime - f.timeWork
						task.LeedTime = leedTime
						f.tasks <- task
					}
				}
			}(i)
		}
		wg.Wait()
	}()
}
