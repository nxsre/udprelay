# udprelay

**udprelay** is a program that accepts UDP connections and relays all incoming data to every known peer. By default, it does this transparently, resending packets to every address it's received a packet from in the last ten minutes; however, with the `-protocol` option, udprelay will implement a lightweight protocol which checks all incoming packets for commands to implement functionality such as channel switching and keep-alives.

For documentation on the program and protocol, check **udprelay**(1) and **udprelay**(7) respectively.

# Building & Installation

A udprelay package for Arch Linux is available in the AUR ([udprelay](https://aur.archlinux.org/packages/udprelay)). If you maintain a udprelay package for your distro, feel free to submit a patch to add a link to this README!

Building udprelay from source only requires an installation of Go and its standard library. Run `make udprelay` to build the udprelay binary.

To make the man pages, you will need to have [scdoc](https://git.sr.ht/~sircmpwn/scdoc/). Simply run `make docs` to generate `udprelay.1` and `udprelay.7`.

# Contributing

Feel free to submit patches or ask questions at [~kt/udprelay@lists.sr.ht](mailto:~kt/udprelay@lists.sr.ht)

