package model

import "time"

type Todolist struct {
	Todothing string    `json:"todothing"`
	Deadline  time.Time `json:"deadline"`
}
type TodolistGET struct {
	Todothing string    `json:"todothing"`
	Deadline  time.Time `json:"deadline"`
	Status    bool      `json:"status"`
	Id        int       `json:"id"`
}
