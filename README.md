# udprelay

**udprelay** is a program that accepts UDP connections and relays all incoming data to every known peer. By default, it does this transparently, resending packets to every address it's received a packet from in the last ten minutes; however, with the `-protocol` option, udprelay will implement a lightweight protocol which checks all incoming packets for commands to implement functionality such as channel switching and keep-alives.

For documentation on the program and protocol, check **udprelay**(1) and **udprelay**(7) respectively.

# Contributing

Feel free to submit patches or ask questions at [~kt/udprelay@lists.sr.ht](mailto:~kt/udprelay@lists.sr.ht)

