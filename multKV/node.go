package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Node struct {
	Port  int
	Peers []string
}

func (n *Node) Start() error {
	address := fmt.Sprintf(":%d", n.Port)
	ln, err := net.Listen("tcp", address)

	
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	
	defer ln.Close()

	log.Printf("Node listening on %s", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("could not accept request: %w", err)
		}

		log.Printf("Accepted connection from %v", conn.RemoteAddr())

		handleRead(conn)
	}
}

func (n *Node) Ping(index int) error {
	if index < 0 || index >= len(n.Peers) {
		return fmt.Errorf("Index: %d passed into ping is invalid", index)
	}

	conn, err := net.DialTimeout("tcp", n.Peers[index], 3 * time.Second)

	if err != nil {
		return fmt.Errorf("Couldn't connect to node: %w", err)
	}
	defer conn.Close()

	time.Sleep(5 * time.Second)

	message := fmt.Sprintf("Hello from node: %d\n", n.Port)
	
	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("Couldn't send message to node: %w", err)
	}

	return nil
}

func read(c net.Conn) error {
	buffer := make([]byte, 1024)

	n, err := c.Read(buffer)
	if err != nil {
		return fmt.Errorf("Read error: %v", err)
	}

	fmt.Printf("Recieved: %s\n", string(buffer[:n]))
	return nil
}

func handleRead(c net.Conn) {
	defer c.Close()

	if err := read(c); err != nil {
		log.Printf("connection error: %v", err)
	}
}