//
// main.go
// go-command-line-todo
//
// Created by Varun Pullur on 05/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package main

import (
	"fmt"
	"go-command-line-todo/internal/cli"
	"go-command-line-todo/internal/store"
	"go-command-line-todo/internal/ui"
	"os"
)

func main() {
	reader := cli.NewReader()

	ui.PrintStoreMenu()
	storeChoice := reader.ReadLine()

	var s store.Store
	switch storeChoice {
	case "1": 
		s = store.NewMemoryStore()
	case "2":
		fmt.Println("File storage not implemented yet.")
		os.Exit(1)
	default:
		fmt.Println("Invalid choice.")
		os.Exit(1)
	}

	app := cli.NewApp(s, reader)
	app.Run()
}