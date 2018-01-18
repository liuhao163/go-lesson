package main

import (
	"net"
	"log"
	"bufio"
	"io"
	"fmt"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":18090");
	if (err != nil) {
		panic(err)
	}
	for {
		conn, err := ln.Accept();
		if (err != nil) {
			log.Fatal("get client connection error: ", err)
		}

		go handleConnection(conn);
	}
}

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		data = strings.TrimSpace(data)
		if (data == "exit") {
			fmt.Println("now close")
			conn.Close()
			break
		}
		fmt.Println(data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}
