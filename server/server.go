package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

// runCommand executes a shell command and returns the output or error as a string.
func runCommand(userInput string) string {
	command := fmt.Sprintf("llm -m Meta-Llama-3-8B-Instruct --option max_tokens 2048 '%s'", userInput)
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing command: %s", err)
	}
	return string(output)
}

// handleConnection handles individual client connections.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		result := runCommand(command)
		fmt.Println(result)
		_, err := conn.Write([]byte(result + "END_OF_RESPONSE\n"))
		if err != nil {
			fmt.Println("Failed to send data to client:", err)
			break
		}
	}
}

// startServer initializes a TCP server that listens for connections on the specified address.
func startServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error listening on %s: %s\n", address, err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Server listening on %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func main() {
	startServer(":12345")
}
