//
// prompts.go
// go-command-line-todo
//
// Created by Varun Pullur on 07/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package ui

import "fmt"

const storeMenuBanner = 
`
=====================================
        Welcome to Todo CLI
=====================================

Choose storage type:
  1. In-Memory (data lost on exit)
  2. File-based (coming soon)

Enter your choice: 
`

const taskMenuFullBanner = `
=====================================
              Todo CLI
=====================================
  1. Add Task
  2. List Tasks
  3. Update Task
  4. Delete Task
  5. Exit
Enter your choice: `

const taskMenuShorBanner = `
  1. Add Task
  2. List Tasks
  3. Update Task
  4. Delete Task
  5. Exit
Enter your choice: `

const invalidChoice = "Invalid Choice"

func PrintStoreMenu() {
	fmt.Print(storeMenuBanner)
}

func PrintTaskMenuFull() {
	fmt.Print(taskMenuFullBanner)
}

func PrintTaskMenuShort() {
	fmt.Print(taskMenuShorBanner)
}

func PrintInvalidChoice() {
	fmt.Print(invalidChoice)
}