//go:build windows

package main

import (
	"os"
	"syscall"

	"golang.org/x/sys/windows"
)

// useUTF8Console switches the console output code page to UTF-8 while the
// process runs, so accented pt-BR messages render correctly in PowerShell and
// cmd. It only acts when stdout is a terminal (not a pipe or file) and
// returns a function that restores the previous code page.
func useUTF8Console() func() {
	var mode uint32
	if err := windows.GetConsoleMode(windows.Handle(os.Stdout.Fd()), &mode); err != nil {
		return func() {}
	}
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getCP := kernel32.NewProc("GetConsoleOutputCP")
	setCP := kernel32.NewProc("SetConsoleOutputCP")
	prev, _, _ := getCP.Call()
	const utf8 = 65001
	if prev == utf8 {
		return func() {}
	}
	if r, _, _ := setCP.Call(uintptr(utf8)); r == 0 {
		return func() {}
	}
	return func() { _, _, _ = setCP.Call(prev) }
}
