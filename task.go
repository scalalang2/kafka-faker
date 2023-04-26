package main

import "time"

type Task func()

type Executor interface {
	Do(task Task)
}

type TaskOnce struct {
	Executor
}

func (t *TaskOnce) Do(task Task) {
	task()
}

type TaskLoop struct {
	Executor
	Delay time.Duration
}

func (t *TaskLoop) Do(task Task) {
	for {
		task()
		time.Sleep(t.Delay)
	}
}
