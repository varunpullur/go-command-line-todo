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
			fmt.Println("Add Task")
		case "2":
			fmt.Println("List All Tasks")
		case "3":
			fmt.Println("Update Task")
		case "4":
			fmt.Println("Delete Task")
		case "5":
			fmt.Println("GoodBye!")
			return
		default:
			ui.PrintInvalidChoice()
		}

		ui.PrintTaskMenuShort()
	}
}