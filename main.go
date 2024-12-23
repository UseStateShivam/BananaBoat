package main

import (
	"fmt"
	"os"
	"os/exec"
)

var runningProcess = make(map[string]*exec.Cmd)

func main() {
	// Check the validity of the input
	if len(os.Args) < 3 {
		fmt.Println("Usage: bananaboat run <command> <args>")
		os.Exit(1)
	}

	// Check if the command is "run", if yes, then start a new process
	if os.Args[1] == "run" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: bananaboat run <command>")
			os.Exit(1)
		}

		// Command to run
		cmd := exec.Command("cmd", "/C", "start", os.Args[2])

		// Attach standard i/o
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running command: %v\n", err)
			os.Exit(1)
		}

		// Store the running process in the map
		runningProcess[os.Args[2]] = cmd

		// Wait for the process to finish
		if err := cmd.Wait(); err != nil {
			fmt.Printf("Error running command: %v\n", err)
			os.Exit(1)
		}
	}
	
	// Check if the command is "stop", if yes, then stop a running process
	if os.Args[1] == "stop" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: bananaboat run <command>")
			os.Exit(1)
		}

		cmdName := os.Args[2]

		// Find the running process
        if cmd, ok := runningProcess[cmdName]; ok {
			if err := cmd.Process.Kill(); err != nil {
				fmt.Printf("Error stopping command: %v\n", err)
				os.Exit(1)
			}
			delete(runningProcess, cmdName)
            fmt.Printf("Command stopped with name '%s'\n", cmdName)
        } else {
			fmt.Printf("No running command for the name '%s'\n", cmdName)
		}
	}
}
