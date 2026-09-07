//
// store.go
// go-command-line-todo
//
// Created by Varun Pullur on 05/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package store

import "go-command-line-todo/internal/todo"

type Store interface {
	Add(t todo.Todo) error
	List() ([]todo.Todo, error)
	Update(t todo.Todo) error
	Delete(id int) error
}