package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.Int("port", 8080, "Type in a port to start up a node (default: 8080)")
	flag.Parse()

	node := Node{
		Port:  *port,
		Peers: []string{"localhost:8002", "localhost:8003"},
	}

	if err := node.Start(); err != nil {
		log.Fatal(err)
	}
}
