package main

import (
	"bytes"
	"log"
	"net"
	"time"
)

var cmdPrefix = []byte("udprelay!")

type Relay struct {
	Log             *log.Logger
	CommandProtocol bool

	Timeout time.Duration

	peers    map[string]*Peer
	channels map[string]*Channel
}

type Channel struct {
	Peers map[string]*Peer
}

type Peer struct {
	Addr    *net.UDPAddr
	Timeout time.Time
	Channel string
}

func (relay *Relay) HandlePacket(conn *net.UDPConn, addr *net.UDPAddr, packet []byte) {
	peer := relay.peers[addr.String()]
	if peer == nil {
		peer = &Peer{
			Addr:    addr,
			Channel: "",
		}
		if relay.peers == nil {
			relay.peers = make(map[string]*Peer)
			relay.channels = make(map[string]*Channel)
			relay.channels[""] = &Channel{
				Peers: make(map[string]*Peer),
			}
		}
		channel := relay.channels[""]
		channel.Peers[peer.Addr.String()] = peer
		relay.peers[addr.String()] = peer
	}
	peer.Timeout = time.Now().Add(relay.Timeout)

	if relay.CommandProtocol {
		if bytes.HasPrefix(packet, cmdPrefix) {
			cmd := packet[len(cmdPrefix):]
			cmd, args := Split2Space(cmd)
			args = TrimSpace(args)
			relay.handleCommand(conn, packet, peer, string(cmd), args)
			return
		}
	}

	sender := peer
	for _, peer := range relay.channels[peer.Channel].Peers {
		if peer.Addr == sender.Addr {
			continue
		}

		if time.Now().After(peer.Timeout) {
			relay.dropPeer(peer)
			continue
		}

		_, err := conn.WriteToUDP(packet, peer.Addr)
		if err != nil {
			relay.Log.Printf("error: writing to %s: %s\n", peer.Addr.String(), err)
			continue
		}
	}
}

func (relay *Relay) dropPeer(peer *Peer) {
	relay.switchChannel(peer, "")
	delete(relay.peers, peer.Addr.String())
}

func (relay *Relay) switchChannel(peer *Peer, channelName string) {
	delete(relay.channels[peer.Channel].Peers, peer.Addr.String())
	if !(peer.Channel == "") && len(relay.channels[peer.Channel].Peers) == 0 {
		delete(relay.channels, peer.Channel)
	}
	peer.Channel = channelName
	channel := relay.channels[channelName]
	if channel == nil {
		channel = &Channel{
			Peers: make(map[string]*Peer),
		}
		relay.channels[channelName] = channel
	}
	channel.Peers[peer.Addr.String()] = peer
}

func (relay *Relay) handleCommand(conn *net.UDPConn, packet []byte, peer *Peer, cmd string, args []byte) {
	switch cmd {
	case "echo":
		_, err := conn.WriteToUDP(packet, peer.Addr)
		if err != nil {
			relay.Log.Printf("error: replying to ping: %s\n", err)
		}
	case "channel":
		relay.switchChannel(peer, string(args))

		_, err := conn.WriteToUDP(packet, peer.Addr)
		if err != nil {
			relay.Log.Printf("error: replying to channel switch message: %s\n", err)
		}
	}
}
