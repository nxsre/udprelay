package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var flagTimeout = flag.String("timeout", "10m", "duration to keep connections alive after their last packet")

type Peer struct {
	Addr    *net.UDPAddr
	Timeout time.Time
}

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [OPTION...] port\n\noptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	timeoutDuration, err := time.ParseDuration(*flagTimeout)
	if err != nil {
		log.Println("error: parsing -timeout: %s\n", err)
		os.Exit(2)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	listenPort, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}

	peers := make(map[string]*Peer, 0)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: listenPort})
	if err != nil {
		panic(err)
	}
	log.Printf("listen: %d\n", listenPort)

	buf := make([]byte, 65536)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("error: %s\n", err)
			continue
		}
		packet := buf[:n]

		peer, exists := peers[addr.String()]
		if !exists {
			log.Printf("connect: %s\n", addr.String())
			peer = &Peer{
				Addr: addr,
			}
			peers[addr.String()] = peer
		}
		peer.Timeout = time.Now().Add(timeoutDuration)

		sender := peer
		for addr, peer := range peers {
			if peer.Addr == sender.Addr {
				continue
			}

			if time.Now().After(peer.Timeout) {
				log.Printf("timeout: %s\n", addr)
				delete(peers, addr)
				continue
			}

			_, err := conn.WriteToUDP(packet, peer.Addr)
			if err != nil {
				log.Printf("error: writing to %s: %s\n", peer.Addr.String(), err)
				continue
			}
		}
	}
}
