//go:build !windows

package main

func useUTF8Console() func() { return func() {} }
