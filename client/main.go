package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: client server_host server_port file_name")
		os.Exit(1)
	}
	serverHost := os.Args[1]
	serverPort := os.Args[2]
	fileName := os.Args[3]

	conn, err := net.Dial("tcp", serverHost+":"+serverPort)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET /%s HTTP/1.1\r\nHost: %s\r\n\r\n", fileName, serverHost)

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			os.Exit(1)
		}
		fmt.Print(line)
		if line == "\r\n" {
			break
		}
	}
	fileContents, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		os.Exit(1)
	}
	fmt.Print(fileContents)
}
