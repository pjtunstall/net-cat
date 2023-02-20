# NET-CAT

1. [THE BRIEF](#1-the-brief)
2. [HOW TO LAUNCH THE SERVER](#2-how-to-launch-the-server)
3. [HOW TO JOIN THE CHAT](#3-how-to-join-the-chat)
4. [HOW TO JOIN ON A DIFFERENT COMPUTER](#4-how-to-join-on-a-different-computer)
5. [A CHOICE WE MADE](#5-system-messages)

## 1. THE BRIEF

The object of this exercise is to write a server in Go that will reproduce some of the functionality of net-cat, a real-time chat application.

## 2. HOW TO LAUNCH THE SERVER

You have two options.

* (Recommended.) Open a terminal and type `go run .`, followed by a port number of your choice, then press ENTER. If no port is specified, the server will listen on 8989 by default.

* Alternatively, to follow the audit questions to the letter, build an excutable file with `go build -o CTPChat main.go` and run it with `./TCPChat $port`, where $port stands for a port number (optional).

Unfortunately no executable can be included in this repository because Gitea won't allow it. Depending on your operating system and the operating system that you want to run it on, you may have to add more flags while building the binary to ensure compatibility. You may also have to grant your operating system permission to run it.

## 3. HOW TO JOIN THE CHAT

To join on the same computer as the server, leave the server running in its terminal and open another terminal. In this other terminal, run the command `nc localhost 8989` to connect as a new client on port 8989, for example. You'll be prompted to enter your name, after which you can join the chat.

## 4. HOW TO JOIN ON A DIFFERENT COMPUTER

First, replace "localhost" in the source code with the IP address of the server. Now build an executable file if you want to run the program that way. Likewise, to connect, you'll need to type `nc` in the client terminal, followed by the IP address of the server and, optionally, the port number, then press ENTER.

To run the server on campus, try to connect to the server via a mobile hotspot. If one combination of computer and phone doesn't work, try another. It can be done. You can't connect over the college wifi network.

To run it from home, you'll need to configure your router settings to allow port forwarding, if your internet provider permits.

Be sure grant access through any intervening firewalls.

## 5. SYSTEM MESSAGES

The example in the instructions suggests that someone (Lee) joining an existing chat (started by Yenlik) sees something like this:

[ENTER YOUR NAME]: Lee
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?

We've chosen to also include "system messages" as it seems useful to let a new client know who's already in the chat, even if they haven't posed any messages yet. Thus:

[ENTER YOUR NAME]: Lee
Yenlik has joined our chat...
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?