//
// app.go
// go-command-line-todo
//
// Created by Varun Pullur on 07/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package cli

import (
	"errors"
	"fmt"
	"go-command-line-todo/internal/commands"
	"go-command-line-todo/internal/store"
	"go-command-line-todo/internal/ui"
)

var (
	ErrorInvalidChoice = errors.New("Invalid choice")
)

type App struct {
	store store.Store
	reader *Reader
}

func NewApp(s store.Store, r *Reader) *App {
	return &App{
		store: s,
		reader: r,
	}
}

func (a *App) Run() {
	ui.PrintTaskMenuFull()

	for {
		choice := a.reader.ReadLine()

		switch choice {
		case "1":
			a.runAddFlow()
		case "2":
			a.runListFlow()
		case "3":
			a.runMarkDoneFlow()
		case "4":
			a.runUpdateFlow()
		case "5":
			a.runDeleteFlow()
		case "6":
			fmt.Println("GoodBye!")
			return
		default:
			ui.PrintInvalidChoice()
		}

		ui.PrintTaskMenuShort()
	}
}

func (a *App) runAddFlow() {
	fmt.Print("Enter task title: ")
	title := a.reader.ReadLine()

	t, err := commands.HandleAdd(title, a.store)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Added task #%d: %s\n", t.ID, t.Title)
}

func (a *App) runListFlow() {
	todos, err := commands.HandleListTodos(a.store)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for _, t := range todos {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Title)
	}
}

func (a *App) runMarkDoneFlow() {
	fmt.Print("Enter task ID to mark done: ")
	id, err := a.reader.ReadInt()
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	t, err := commands.HandleMarkAsDone(id, a.store)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Marked task #%d as done.\n", t.ID)
}

func (a *App) runUpdateFlow() {
	fmt.Print("Enter task ID to update: ")
	id, err := a.reader.ReadInt()
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	fmt.Print("Enter new title: ")
	newTitle := a.reader.ReadLine()

	t, err := commands.HandleUpdate(id, newTitle, a.store)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Updated task #%d: %s\n", t.ID, t.Title)
}

func (a *App) runDeleteFlow() {
	fmt.Print("Enter task ID to delete: ")
	id, err := a.reader.ReadInt()
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	err = commands.HandleDelete(id, a.store)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deleted task #%d.\n", id)
}