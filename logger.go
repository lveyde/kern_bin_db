/*
 * Copyright (c) 2023 Red Hat, Inc.
 * SPDX-License-Identifier: GPL-2.0-or-later
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	debugEnabled = false

	InfoLogger  *log.Logger
	DebugLogger *log.Logger
)

func init() {
	logFile, err := os.OpenFile(filepath.Base(os.Args[0])+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	DebugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime)
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
}

func Debug(v ...any) {
	if debugEnabled {
		DebugLogger.Println(v...)
	}
}

func Info(v ...any) {
	InfoLogger.Println(v...)
}

func FormatStruct(v any) string {
	b, err := json.MarshalIndent(v, "=", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(string(b))
}
