package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Custom writer to count the bytes written.
type ByteCounter struct {
	Count int
}

// Write method for the ByteCounter.
func (bc *ByteCounter) Write(p []byte) (int, error) {
	n := len(p)
	bc.Count += n
	return n, nil
}

func handleCommand(command string) {
	// print the command
	fmt.Printf("> `%s`\n", command)

	// check if the user wants to execute the command
	fmt.Print("execute this command? (y/n): ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal("Error reading input: ", err)
	}

	if input != "y" {
		fmt.Println("exiting...")
		os.Exit(0)
	}

	fmt.Println()

	// initialize the byte counter
	byteCounter := &ByteCounter{}
	multiWriter := io.MultiWriter(os.Stdout, byteCounter)

	// execute the command
	cmd := exec.Command("zsh", "-c", command)
	cmd.Stdout = multiWriter
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println("⚠️ ", strings.Split(command, " ")[0], "exited with error: ", err, "\n")
	}

	// if the byte counter is not 0, print a newline
	if byteCounter.Count != 0 {
		fmt.Println()
	}

	fmt.Println("---")
}
