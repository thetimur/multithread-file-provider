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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
		if line == "\r\n" {
			break
		}
	}
}
