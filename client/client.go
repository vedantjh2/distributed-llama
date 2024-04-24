package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Enter commands to execute or 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}

		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Println("Failed to send command:", err)
			break
		}

		// Read and display the complete response from the server
		responseReader := bufio.NewReader(conn)
		fmt.Println("Response from server: ")
		for {
			responseLine, err := responseReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break // End of file (or stream) means we're done reading.
				}
				fmt.Println("Failed to read response:", err)
				break
			}
			// Check if this is the end of the response
			if responseLine == "END_OF_RESPONSE\n" {
				break
			}
			fmt.Print(responseLine)
		}
	}
}
