//
// commands.go
// go-command-line-todo
//
// Created by Varun Pullur on 08/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package commands

import (
	"errors"
	"go-command-line-todo/internal/store"
	"go-command-line-todo/internal/todo"
	"strings"
	"time"
)

var (
	ErrEmptyTitle = errors.New("Title of todo cannot be empty")
	ErrTodoNotFound = errors.New("Todo not found")
)

func HandleAdd(title string, s store.Store) (todo.Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return todo.Todo{}, ErrEmptyTitle
	}

	t := todo.Todo{Title: title, CreatedAt: time.Now()}
	return s.Add(t)
}

func HandleListTodos(s store.Store) ([]todo.Todo, error) {
	return s.List()
}

func HandleMarkAsDone(id int, s store.Store) (todo.Todo, error) {
	todos, err := s.List()

	if err != nil {
		return todo.Todo{}, err
	}

	var t todo.Todo
	found := false
	for i := range todos {
		if todos[i].ID == id {
			t = todos[i]
			found = true
			break
		}
	}
	if !found {
		return todo.Todo{}, ErrTodoNotFound
	}

	t.MarkAsDone()

	err = s.Update(t)
	if err != nil {
		return todo.Todo{}, err
	}

	return t, nil
}

func HandleUpdate(id int, newTitle string, s store.Store) (todo.Todo, error) {
	todos, err := s.List()

	if err != nil {
		return todo.Todo{}, err
	}

	var t todo.Todo
	found := false
	for i := range todos {
		if todos[i].ID == id {
			t = todos[i]
			found = true
			break
		}
	}
	if !found {
		return todo.Todo{}, ErrTodoNotFound
	}

	t.Title = newTitle
	err = s.Update(t)

	return t, err
}

func HandleDelete(id int, s store.Store) error {
	return s.Delete(id)
}