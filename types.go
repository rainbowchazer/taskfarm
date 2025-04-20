package main

type Task struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Param  int    `json:"param"`
	Result int    `json:"result"`
	Status string `json:"status"`
}
