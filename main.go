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
var flagProtocol = flag.Bool("protocol", false, "enables the udprelay command protocol")

var flagVersion = flag.Bool("version", false, "print the version to stdout and exit immediately")

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "udprelay %s\nusage: %s [OPTION...] port\n\noptions:\n", Version, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "\nsee `man udprelay.1` for more information")
	}
	flag.Parse()

	if *flagVersion {
		fmt.Printf("udprelay %s\n", Version)
		os.Exit(0)
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

	log.Printf("udprelay %s\n\n", Version)

	timeoutDuration, err := time.ParseDuration(*flagTimeout)
	if err != nil {
		log.Println("error: parsing -timeout: %s\n", err)
		os.Exit(2)
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: listenPort})
	if err != nil {
		panic(err)
	}
	log.Printf("listen: %d\n", listenPort)

	relay := &Relay{
		Log:             log.New(os.Stderr, "", 0),
		Timeout:         timeoutDuration,
		CommandProtocol: *flagProtocol,
	}

	buf := make([]byte, 65536)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("error: %s\n", err)
			continue
		}
		packet := buf[:n]

		relay.HandlePacket(conn, addr, packet)
	}
}
