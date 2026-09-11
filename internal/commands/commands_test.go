//
// commands_test.go
// go-command-line-todo
//
// Created by Varun Pullur on 11/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package commands

import (
	"errors"
	"go-command-line-todo/internal/store"
	"go-command-line-todo/internal/todo"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestHandleAdd(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantErr   error
		wantTitle string
	}{
		{
			name: "valid title is added",
		 	title: "Buy milk",
			wantTitle: "Buy milk",
		},
		{
			name: "title with surrounding whitespace is trimmed",
			title: "  Walk dog  ",
			wantTitle: "Walk dog",
		},
		{
			name: "empty title is rejected",
			title: "", 
			wantErr: ErrEmptyTitle,
		},
		{
			name: "whitespace-only title is rejected", 
			title: "   ", 
			wantErr: ErrEmptyTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := store.NewMockStore(ctrl)

			if tt.wantErr == nil {
				mockStore.EXPECT().
					Add(gomock.Any()).
					DoAndReturn(func(got todo.Todo) (todo.Todo, error) {
						if got.Title != tt.wantTitle {
							t.Errorf("Add called with title %q, want %q", got.Title, tt.wantTitle)
						}
						got.ID = 1
						return got, nil
					})
			} else {
				mockStore.EXPECT().Add(gomock.Any()).Times(0)
			}

			got, err := HandleAdd(tt.title, mockStore)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Title != tt.wantTitle {
				t.Errorf("got title %q, want %q", got.Title, tt.wantTitle)
			}
		})
	}

	t.Run("propagates store error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)
		wantErr := errors.New("store failure")

		mockStore.EXPECT().Add(gomock.Any()).Return(todo.Todo{}, wantErr)

		_, err := HandleAdd("Buy milk", mockStore)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})
}

func TestHandleListTodos(t *testing.T) {
	t.Run("returns todos from store", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)

		want := []todo.Todo{{ID: 1, Title: "Buy milk"}, {ID: 2, Title: "Walk dog"}}
		mockStore.EXPECT().List().Return(want, nil)

		got, err := HandleListTodos(mockStore)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got) != len(want) {
			t.Errorf("got %d todos, want %d", len(got), len(want))
		}
	})

	t.Run("propagates store error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)
		wantErr := errors.New("list failure")

		mockStore.EXPECT().List().Return(nil, wantErr)

		_, err := HandleListTodos(mockStore)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})
}

func TestHandleMarkAsDone(t *testing.T) {
	tests := []struct {
		name      string
		id        int
		wantErr   error
		setupMock func(m *store.MockStore)
	}{
		{
			name: "marks existing todo as done",
			id:   1,
			setupMock: func(m *store.MockStore) {
				existing := []todo.Todo{{ID: 1, Title: "Buy milk"}}
				m.EXPECT().List().Return(existing, nil)
				m.EXPECT().
					Update(gomock.Any()).
					DoAndReturn(func(got todo.Todo) error {
						if !got.Done {
							t.Errorf("expected Update to be called with Completed=true")
						}
						return nil
					})
			},
		},
		{
			name:    "id not found returns ErrTodoNotFound",
			id:      999,
			wantErr: ErrTodoNotFound,
			setupMock: func(m *store.MockStore) {
				existing := []todo.Todo{{ID: 1, Title: "Buy milk"}}
				m.EXPECT().List().Return(existing, nil)
				m.EXPECT().Update(gomock.Any()).Times(0)
			},
		},
		{
			name:    "empty store returns ErrTodoNotFound",
			id:      1,
			wantErr: ErrTodoNotFound,
			setupMock: func(m *store.MockStore) {
				m.EXPECT().List().Return([]todo.Todo{}, nil)
				m.EXPECT().Update(gomock.Any()).Times(0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := store.NewMockStore(ctrl)
			tt.setupMock(mockStore)

			got, err := HandleMarkAsDone(tt.id, mockStore)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !got.Done {
				t.Errorf("expected returned todo to be marked done, got %+v", got)
			}
		})
	}

	t.Run("propagates List error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)
		wantErr := errors.New("list failure")

		mockStore.EXPECT().List().Return(nil, wantErr)

		_, err := HandleMarkAsDone(1, mockStore)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})

	t.Run("propagates Update error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)
		wantErr := errors.New("update failure")

		mockStore.EXPECT().List().Return([]todo.Todo{{ID: 1, Title: "Buy milk"}}, nil)
		mockStore.EXPECT().Update(gomock.Any()).Return(wantErr)

		_, err := HandleMarkAsDone(1, mockStore)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})
}

func TestHandleUpdate(t *testing.T) {
	tests := []struct {
		name      string
		id        int
		newTitle  string
		wantErr   error
		wantTitle string
		setupMock func(m *store.MockStore)
	}{
		{
			name:      "updates title of existing todo",
			id:        1,
			newTitle:  "Buy oat milk",
			wantTitle: "Buy oat milk",
			setupMock: func(m *store.MockStore) {
				existing := []todo.Todo{{ID: 1, Title: "Buy milk"}}
				m.EXPECT().List().Return(existing, nil)
				m.EXPECT().
					Update(gomock.Eq(todo.Todo{ID: 1, Title: "Buy oat milk"})).
					Return(nil)
			},
		},
		{
			name:     "id not found returns ErrTodoNotFound",
			id:       999,
			newTitle: "doesn't matter",
			wantErr:  ErrTodoNotFound,
			setupMock: func(m *store.MockStore) {
				existing := []todo.Todo{{ID: 1, Title: "Buy milk"}}
				m.EXPECT().List().Return(existing, nil)
				m.EXPECT().Update(gomock.Any()).Times(0)
			},
		},
		{
			name:     "empty store returns ErrTodoNotFound",
			id:       1,
			newTitle: "doesn't matter",
			wantErr:  ErrTodoNotFound,
			setupMock: func(m *store.MockStore) {
				m.EXPECT().List().Return([]todo.Todo{}, nil)
				m.EXPECT().Update(gomock.Any()).Times(0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := store.NewMockStore(ctrl)
			tt.setupMock(mockStore)

			got, err := HandleUpdate(tt.id, tt.newTitle, mockStore)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Title != tt.wantTitle {
				t.Errorf("got title %q, want %q", got.Title, tt.wantTitle)
			}
		})
	}

	t.Run("propagates List error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)
		wantErr := errors.New("list failure")

		mockStore.EXPECT().List().Return(nil, wantErr)

		_, err := HandleUpdate(1, "new title", mockStore)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})

	t.Run("propagates Update error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := store.NewMockStore(ctrl)
		wantErr := errors.New("update failure")

		mockStore.EXPECT().List().Return([]todo.Todo{{ID: 1, Title: "Buy milk"}}, nil)
		mockStore.EXPECT().Update(gomock.Any()).Return(wantErr)

		_, err := HandleUpdate(1, "new title", mockStore)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})
}

func TestHandleDelete(t *testing.T) {
	tests := []struct {
		name      string
		deleteErr error
		wantErr   bool
	}{
		{"passes through success", nil, false},
		{"passes through store error", errors.New("delete failure"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := store.NewMockStore(ctrl)

			mockStore.EXPECT().Delete(1).Return(tt.deleteErr)

			err := HandleDelete(1, mockStore)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}