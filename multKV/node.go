package main

import (
	"fmt"
	"log"
	"net"
)

func startNode(port int) {
	address := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", address)

	
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
	
	defer ln.Close()

	log.Printf("Node listening on %s", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("could not accept request: %v", err)
			continue
		}	

		log.Printf("Accepted connection from %v", conn.RemoteAddr())

		go func(c net.Conn) {
			defer c.Close()
			

			log.Printf("Sending response to %v", c.RemoteAddr())
		} (conn)
	}
}