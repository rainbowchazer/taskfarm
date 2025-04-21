package main

import (
	"log"
	"net/http"
)

func main() {
	startWorker("worker-1")
	startWorker("worker-2")
	startWorker("worker-3")

	http.HandleFunc("/submit-task", submitTaskHandler)
	http.HandleFunc("/get-task", getTaskHandler)
	http.HandleFunc("/tasks", listTasksHandler)
	http.HandleFunc("/metrics", metricHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
