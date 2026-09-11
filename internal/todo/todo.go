//
// todo.go
// go-command-line-todo
//
// Created by Varun Pullur on 05/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package todo

import "time"

type Todo struct {
	ID int
	Title string
	Done bool
	CreatedAt time.Time
}

func (t *Todo) MarkAsDone() {
	t.Done = true
}