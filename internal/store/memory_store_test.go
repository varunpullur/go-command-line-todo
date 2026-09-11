//
// memory_store_tests.go
// go-command-line-todo
//
// Created by Varun Pullur on 08/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package store

import (
	"errors"
	"go-command-line-todo/internal/todo"
	"testing"
)

func TestMemoryStore_Add(t *testing.T) {
	tests := []struct {
		name string
		title string
		wantID int
	} {
		{"first task gets ID 1", "Buy milk", 1},
        {"second task gets ID 2", "Walk dog", 2},
	}

	s := NewMemoryStore()

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			got, err := s.Add(todo.Todo{Title: tt.title})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.ID != tt.wantID {
				t.Errorf("got ID %d, want %d", got.ID, tt.wantID)
			}
		})
	}
}

func TestMemoryStore_List(t *testing.T) {
	tests := []struct {
		name string
		titles []string
		wantCount int
	} {
		{"Empty store returns nil", nil, 0},
		{"Single Todo", []string{"Buy Milk"}, 1},
		{"Multiple Todos", []string{"Buy milk", "walk dog", "pay bills"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := NewMemoryStore()
			for _, title := range tt.titles {
				s.Add(todo.Todo{Title: title})
			}

			got, err := s.List()

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(got) != tt.wantCount {
				t.Errorf("got %d todos, want %d", len(got), tt.wantCount)
			}
		})
	}
}

func TestMemoryStore_Update(t *testing.T) {
	tests := []struct {
		name string
		titles []string
		updateID int
		newTitle string
		wantErr error
		wantTitle string
	} {
		{
			name: "empty store returns ErrEmptyTodos",
			titles: nil,
			updateID: 1,
			newTitle: "Doesn't Matter",
			wantErr: ErrEmptyTodos,
		},
		{
			name: "id not found returns ErrTodoNotFound",
			titles: []string{"Buy Milk"},
			updateID: 999,
			newTitle: "doesn't matter",
			wantErr: ErrTodoNotFound,
		},
		{
			name: "existing id gets updated",
			titles: []string{"Buy milk"},
			updateID: 1,
			newTitle: "Buy oat milk",
			wantErr: nil,
			wantTitle: "Buy oat milk",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMemoryStore()
			for _, title := range tt.titles {
				s.Add(todo.Todo{Title: title})
			}

			err := s.Update(todo.Todo{ID: tt.updateID, Title: tt.newTitle})

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			list, _ := s.List()
			for _, item := range list {
				if item.ID == tt.updateID && item.Title != tt.wantTitle {
					t.Errorf("got title %q, want %q", item.Title, tt.wantTitle)
				}
			}
		})
	}
}

func TestMemoryStore_Delete(t *testing.T) {
	tests := []struct {
		name string
		titles []string
		deleteID int
		wantErr error
		wantCount int
	} {
		{
			name: "empty store returns ErrEmptyTodos",
			titles: nil,
			deleteID: 1,
			wantErr: ErrEmptyTodos,
		},
		{
			name: "id not found returns ErrTodoNotFound",
			titles: []string{"Buy Milk"},
			deleteID: 999,
			wantErr: ErrTodoNotFound,
		},
		{
			name: "deletes existing ID",
			titles: []string{"Buy Milk", "walk dog"},
			deleteID: 1,
			wantErr: nil,
			wantCount: 1,
		},
		{
			name: "deletes last remaining todo",
			titles: []string{"Buy Milk"},
			deleteID: 1,
			wantErr: nil,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMemoryStore()
			for _, title := range tt.titles {
				s.Add(todo.Todo{Title: title})
			}

			err := s.Delete(tt.deleteID)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			list, _ := s.List() 
			if len(list) != tt.wantCount {
				t.Errorf("got %d todos remaining, want %d", len(list), tt.wantCount)
			}

			for _, item := range list {
				if item.ID == tt.deleteID {
					t.Errorf("Deleted todo with ID %d is still present", tt.deleteID)
				}
			}
		})
	}
}