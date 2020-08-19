udprelay(1) "udprelay v1.0.1"

# NAME

udprelay \- server to relay UDP connections to connected peers.

# SYNOPSIS

*udprelay* [_option_...] _port_

# DESCRIPTION

udprelay accepts UDP connections on _port_ and relays all incoming data to every
peer which it has received data from within a certain timeframe defined by the
*-timeout* option. When it hasn't received a packet from a peer in that time, it
quietly drops all internal state related to that peer and stops relaying packets
to it until it receives a packet from it again.

udprelay also features an optional command protocol that allows for more
advanced functionality such as channels. This protocol is documented in
*udprelay*(7) and may be enabled by passing the *-protocol* flag.

# OPTIONS

*-protocol*
	Enables the udprelay command protocol. When enabled, packets beginning
	with the string *udprelay!* will not be relayed and will instead be
	handled by udprelay in accordance with the protocol defined in
	*udprelay*(7). When unset, all packets will always be relayed.

*-timeout* _duration_
	Set the amount of time to leave connections open without receiving any
	packets from a peer. _duration_ is a sequence of decimal numbers with
	unit suffixes, such as *10m*, *120s*, and *5m48s*.

*-version*
	Prints the version of udprelay to stdout and exits with a code of 0.

# SEE ALSO

*udprelay*(7)
