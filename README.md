# NET-CAT

1. [THE BRIEF](#1-the-brief)
2. [HOW TO LAUNCH THE SERVER](#2-how-to-launch-the-server)
3. [HOW TO JOIN THE CHAT](#3-how-to-join-the-chat)
4. [HOW TO JOIN ON A DIFFERENT COMPUTER](#4-how-to-join-on-a-different-computer)
5. [SYSTEM MESSAGES](#5-system-messages)

## 1. THE BRIEF

The object of this exercise is to write a server in Go that will reproduce some of the functionality of net-cat, a real-time chat application. Specifically, our program should substitute for the `-l` option of net-cat. It should listen on a chosen port for clients launched by the `nc` command and allow them to chat in real time, using TCP and goroutines.

## 2. HOW TO LAUNCH THE SERVER

You have two options.

* (Recommended.) Open a terminal and type `go run .`, followed by a port number of your choice, then press ENTER. If no port is specified, the server will listen on 8989 by default.

* Alternatively, to follow the audit questions to the letter, build an excutable file with `go build -o CTPChat main.go` and run it with `./TCPChat $port`, where $port stands for a port number (optional).

Unfortunately no executable can be included in this repository because Gitea won't allow it. If you want to run the executable on a different operating system from the one it was built on, you'll have to set certain environment variables before you build. This is to ensure compatibility. You can run `go env GOOS GOARCH` to find out your operating system and architecture. You can find a full list of supported values for GOOS and GOARCH in the Go documentation. Say you want to build the executable for macOS with architecture x64. In that case, you'd run `env GOOS=darwin GOARCH=amd64 go build`.

## 3. HOW TO JOIN THE CHAT

To join on the same computer as the server, leave the server running in its terminal and open another terminal. In this other terminal, run the command `nc localhost 8989` to connect as a new client on port 8989, for example. You'll be prompted to enter your name, after which you can join the chat.

## 4. HOW TO JOIN ON A DIFFERENT COMPUTER

First, replace "localhost" in the source code with the IP address of the server.

Next, build an executable file if you want to run the program that way. To connect, you'll need to open a terminal and type `nc`, followed by the IP address of the server and, optionally, the port number, then press ENTER.

If you want to run the server on campus, be warned that we can't connect over the college wifi network. For reasons of security, we aren't permitted. You could try to connect to the server via a mobile hotspot. If one combination of computer and phone doesn't work, try another. It can be done.

To run it from home, you'll need to configure your router settings to allow port forwarding, if your internet provider permits.

For greater stability, a free dynamic domain name service, such as DuckDNS or No-IP, could be used to associate a domain name with your IP address, even if your IP address is dynamically assigned by your Internet Service Provider (ISP) and changes periodically.

A VPN might be used to bypass restrictions and connect over a local network. Third-party hosting services offer another option for different devices to connect.

Be sure to grant access through any intervening firewalls.

## 5. SYSTEM MESSAGES

The example in the instructions suggests that someone (Lee) joining an existing chat (started by Yenlik) sees something like this:

```
[ENTER YOUR NAME]: Lee
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?
```

We've chosen to also include "system messages" as it seems useful to let a new client know who's already in the chat, even if they haven't posed any messages yet. Thus:

```
[ENTER YOUR NAME]: Lee
Yenlik has joined our chat...
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?
```