//
// reader.go
// go-command-line-todo
//
// Created by Varun Pullur on 07/09/26.
// Copyright © 2026 Varun Pullur. All rights reserved.
//

package cli

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader() *Reader {
	return &Reader{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (r *Reader) ReadLine() string {
	if !r.scanner.Scan() {
		return ""
	}

	return strings.TrimSpace(r.scanner.Text())
}

func (r *Reader) ReadInt() (int, error) {
	line := r.ReadLine()
	return strconv.Atoi(line)
}