package web

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(arg_f func() error) *Task {
	t := Task{
		f: arg_f,
	}
	return &t
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
	jobsChannel  chan *Task
	workerNum    int
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		jobsChannel:  make(chan *Task),
		workerNum:    cap,
	}
	return &p
}

func (p *Pool) worker(workerId int) {
	for task := range p.jobsChannel {
		task.Execute()
		fmt.Println("worker Id:", workerId, "执行了一个任务")
	}
}

func (p *Pool) run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
	for task := range p.EntryChannel {
		p.jobsChannel <- task

	}
}

func Run() {
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	p := NewPool(8)
	count := 0
	go func() {
		for {
			p.EntryChannel <- t
			count++
			time.Sleep(time.Second)
			fmt.Println("task Id:", count)
		}
	}()
	p.run()
}
