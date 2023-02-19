// Run the following command to build an executable file
// if you want to test it exeactly as the audit questions
// suggest. You may have to add flags to indicate which
// OS you want to run it on. Look them up or ask ChatGPT.
// You may also have to run commands to get permission to
// run the executable file.

// Apparently we're not allowed to push executable files
// to Gitea.

// go build -o CTPChat main.go

// Alternatively, if this is a faff, you can just run
// `go run main.go`  or `go run .` while in the same
// directory as this source code file.

// In that case, where the audit questions say to run
// `./TCPChat`, you type `go run .` and where they say
// run `./TCPChat 2525`, you type `go run . 2525`.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	host      = "localhost"
	port      = 8989
	chat      = []string{}
	clients   = map[net.Conn]bool{}
	mu        = new(sync.Mutex)
	linuxLogo = fmt.Sprint(
		"Welcome to TCP-Chat!\n",
		"  _nnnn_\n",
		" dGGGGMMb\n",
		"@p~qp~~qMb\n",
		"M|@||@) M|\n",
		"@,----.JM|\n",
		"JS^\\__/  qKL\n",
		"dZP        qKRb\n",
		"dZP          qKKb\n",
		"fZP            SMMb\n",
		"HZM            MMMM\n",
		"FqM            MMMM\n",
		"__| \".        |\\dS\"qML\n",
		"|    `.       | `' \\Zq\n",
		"_)      \\.___.,|     .'\n",
		"\\____   )MMMMMP|   .'\n",
		"     `-'       `--'\"`\n",
	)
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	startServer()
}

func startServer() {
	if len(os.Args) == 2 {
		chosenPort, err := atoi(os.Args[1])
		if err != nil {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}
		port = chosenPort
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	log.Printf("Listening for connections on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			mu.Lock()
			log.Printf("Error accepting connection from client: %s", err)
			mu.Unlock()
		} else {
			go processClient(conn)
		}
	}
}

func enterName(conn net.Conn) (string, error) {
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		mu.Lock()
		log.Printf("Error reading client name: %s", err)
		mu.Unlock()
		return "", err
	}
	if name == "\n" {
		name, err = enterName(conn)
	}
	name = strings.TrimRight(name, "\n")
	return name, err
}

func processClient(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(linuxLogo))
	if len(clients) > 9 {
		conn.Write([]byte("TCPChat is full right now. Please try again later.\n"))
		return
	}
	name, err := enterName(conn)
	if err != nil {
		return
	}
	mu.Lock()
	clients[conn] = true
	chat = append(chat, fmt.Sprintf("%s has joined our chat...\n", name))
	fmt.Print(chat[len(chat)-1])
	for c := range clients {
		if clients[c] && c != conn {
			c.Write([]byte(chat[len(chat)-1]))
		}
	}
	for msg := 0; msg+1 < len(chat); msg++ {
		conn.Write([]byte(chat[msg]))
	}
	conn.Write([]byte(">"))
	mu.Unlock()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			mu.Lock()
			chat = append(chat, fmt.Sprintf("%s has left our chat...\n", name))
			fmt.Print(chat[len(chat)-1])
			for c := range clients {
				c.Write([]byte(chat[len(chat)-1]))
			}
			mu.Unlock()
			return
		}
		timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
		messageWithTimestamp := fmt.Sprintf("[%s][%s]%s", timestamp, name, message)
		conn.Write([]byte("\033[1A\033[2K\r"))
		mu.Lock()
		if message != "\n" {
			fmt.Print(messageWithTimestamp)
			chat = append(chat, messageWithTimestamp)
			// 	\033 is the escape character, which signals the start of a control character sequence
			// [1A moves the cursor up one line
			// [2K clears the entire line
			// \r moves the cursor back to the beginning of the line
			for c := range clients {
				c.Write([]byte(chat[len(chat)-1]))
			}
		} else {
			conn.Write([]byte(messageWithTimestamp))
		}
		conn.Write([]byte(">"))
		mu.Unlock()
	}
}

func atoi(s string) (int, error) {
	var result int
	var err error
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid input string: %s", s)
		}
		result = result*10 + int(c-'0')
	}
	return result, err
}
