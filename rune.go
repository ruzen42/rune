package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [args...]\n", os.Args[0])
		os.Exit(1)
	}

	if err := checkUID(); err != nil {
		fmt.Fprintf(os.Stderr, "access denied: %v\n", err)
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdPath, err := exec.LookPath(cmdName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "command not found: %s\n", cmdName)
		os.Exit(1)
	}

	if err := syscall.Setgid(0); err != nil {
		fmt.Fprintln(os.Stderr, "setgid failed:", err)
		os.Exit(1)
	}

	if err := syscall.Setuid(0); err != nil {
		fmt.Fprintln(os.Stderr, "setuid failed:", err)
		os.Exit(1)
	}

	safeEnv := []string{
		"PATH=" + os.Getenv("PATH"),
		"HOME=/root",
		"USER=root",
		"LOGNAME=root",
		"SHELL=/bin/sh",
		"TERM=" + os.Getenv("TERM"),
	}

	err = syscall.Exec(cmdPath, os.Args[1:], safeEnv)
	if err != nil {
		fmt.Fprintln(os.Stderr, "exec failed:", err)
		os.Exit(1)
	}
}

func checkUID() error {
	data, err := os.ReadFile("/etc/runeid")
	if err != nil {
		return fmt.Errorf("cannot read /etc/runeid: %w", err)
	}

	targetUID, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return fmt.Errorf("invalid UID format in /etc/runeid: %w", err)
	}

	if os.Getuid() != targetUID {
		return fmt.Errorf("operation not permitted")
	}

	return nil
}
