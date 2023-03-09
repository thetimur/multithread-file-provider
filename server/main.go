package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: server concurrency_level")
		os.Exit(1)
	}

	concurrencyLevel, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid concurrency level:", err)
		os.Exit(1)
	}

	semaphore := make(chan bool, concurrencyLevel)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleRequest(conn, semaphore)
	}
}

func handleRequest(conn net.Conn, semaphore chan bool) {
	semaphore <- true
	defer func() {
		<-semaphore
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	requestString := scanner.Text()

	file := strings.Split(requestString, " ")[1]

	fileContents, err := ioutil.ReadFile(file[1:])

	if err != nil {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte(fmt.Sprintf("Content-Length: %d\r\n", len(fileContents))))
	conn.Write([]byte("Content-Type: text/plain\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write(fileContents)
	conn.Write([]byte("\r\n"))
}
