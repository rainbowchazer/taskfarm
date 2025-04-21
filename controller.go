package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var (
	taskQueue   = make(chan *Task, 100)
	taskResults = make(map[string]*Task)
	mutex       = sync.Mutex{}
)

func submitTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	task.ID = uuid.New().String()
	task.Status = "pending"

	mutex.Lock()
	taskResults[task.ID] = &task
	mutex.Unlock()

	taskQueue <- &task

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mutex.Lock()
	defer mutex.Unlock()

	task, ok := taskResults[id]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var tasks []*Task
	for _, t := range taskResults {
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	metrics := map[string]int{
		"completed_tasks": completedTasks,
		"queue_length":    len(taskQueue),
	}

	json.NewEncoder(w).Encode(metrics)
}
