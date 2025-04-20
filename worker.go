package main

import (
	"fmt"
	"time"
)

func startWorker(name string) {
	go func() {
		for task := range taskQueue {
			fmt.Printf("[%s]Взял задачу %s (%s %d)\n", name, task.ID, task.Type, task.Param)

			switch task.Type {
			case "fib":
				task.Result = fib(task.Param)
			case "square":
				task.Result = task.Param * task.Param
			case "double":
				task.Result = task.Param * 2
			default:
				fmt.Printf("[%s]Неизвестный тип задачи:%s\n", name, task.Type)
				task.Status = "error"
				continue
			}

			task.Status = "done"

			fmt.Printf("[%s]Выполнил задачу %s: %d\n", name, task.ID, task.Result)
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
