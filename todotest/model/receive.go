package model

import "time"

type Todolist struct {
	Todothing string    `json:"todothing"`
	Deadline  time.Time `json:"deadline"`
}
