udprelay(7)

# NAME

udprelay \- protocol for additional functionality implemented by *udprelay*(1)

# COMMANDS

Packets beginning with the string _udprelay!_ are interpreted as commands and are not forwarded to other peers. Each commands has a _head_ and a _body_. The _head_ is a sequence of lowercase alphabetical characters identifying the command to be executed. The _body_ is information passed to the command. The _body_ may contain any characters, but any spacing characters (one of 0x09, 0x0A, 0x0B, 0x0C, 0x0D, or 0x20) at the beginning or end of the body will be trimmed and ignored. The _head_ and _body_ are separated by one of the aforementioned spacing characters. If the _body_ is empty, you may omit the space and only send the _head_. The following are examples of valid command calls:

	udprelay!channel Example

	udprelay!echo

	udprelay!echo Hello, world.

A packet starting with the string _udprelay!_ will never be relayed to other peers. If a call is made to a command that doesn't exist, it is ignored.

## channel

The *channel* command switches the peer to a different channel identified by the first string of non-spacing-characters in the command's _body_. If the _body_ is empty, the peer is switched to the default unnamed channel. Peers will only receive and send messages sent by and to peers who have connected to the same channel. After switching the peer's channel, the server will send the command's packet back to its originator unchanged.

## echo

Upon receiving an *echo* command, udprelay will resend the command's packet to its originator unchanged. Anything or nothing may be passed as arguments to this command; it will not interpret any of the remaining data in the packet.

# SEE ALSO

*udprelay*(1)
