//
// memory_store.go
// go-command-line-todo
//
// Created by Varun Pullur on 05/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package store

import (
	"errors"
	"go-command-line-todo/internal/todo"
	"slices"
)

var (
	ErrEmptyTodos = errors.New("Store is empty.")
	ErrTodoNotFound = errors.New("Todo Not Found in store")
)

type MemoryStore struct {
	todos []todo.Todo
	nextId int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos: []todo.Todo{},
		nextId: 1,
	}
}

func (m *MemoryStore) Add(t todo.Todo) error {
	t.ID = m.nextId
	m.nextId++

	m.todos[t.ID] = t
	return nil
}

func (m *MemoryStore) List() ([]todo.Todo, error) {
	if len(m.todos) == 0 {
		return nil, ErrEmptyTodos
	} 

	return m.todos, nil
}

func (m *MemoryStore) Update(t todo.Todo) error {
	if len(m.todos) == 0 {
		return ErrEmptyTodos
	} 

	for i := range m.todos {
		if t.ID == m.todos[i].ID {
			m.todos[i] = t
			return nil
		}
	}

	return ErrTodoNotFound
}

func (m *MemoryStore) Delete(id int) error {
	if len(m.todos) == 0 {
		return ErrEmptyTodos
	} 

	for i := range m.todos {
		if m.todos[i].ID == id {
			m.todos = slices.Delete(m.todos, i, i+1)
			return nil
		}
	}

	return ErrTodoNotFound
}