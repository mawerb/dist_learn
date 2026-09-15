package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func pingLoop(node *Node) {
	if len(node.Peers) == 0 {
		return
	}

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		for i := range node.Peers {
			fmt.Printf(
				"Node %d: pinging node %s\n",
				node.Port,
				node.Peers[i],
			)
			
			if err := node.Ping(i); err != nil {
				log.Printf("Ping failed: %v", err)
			}
		}
	}
}

func main() {
	port := flag.Int("port", 8080, "Type in a port to start up a node (default: 8080)")
	peerList := flag.String(
		"peers",
		"",
		"comma-separated peer addresses",
	)
	flag.Parse()

	var peers []string
	if (*peerList != "") {
		peers = strings.Split(*peerList, ",")
	}

	node := Node{
		Port:  *port,
		Peers: peers,
	}

	errCh := make(chan error)

	go func() {
		errCh <- node.Start()
	}()

	pingLoop(&node)
}
