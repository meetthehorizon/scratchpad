package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	util "github.com/meetthehorizon/scratchpad/tcp-echo-server/internal"
)

func main() {
	serverAddr := util.GetServerAddr()

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("connected to server: %v\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("what do you want to say to the server?")
	if !scanner.Scan() {
		return // no input (Ctrl-D)
	}
	msg := scanner.Text()

	// Send the message. The server's ReadString('\n') waits for a newline,
	// so we must terminate the message with one.
	if _, err := fmt.Fprintf(conn, "%s\n", msg); err != nil {
		fmt.Println("error sending:", err)
		return
	}

	// Read the server's reply (uppercased echo).
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("error reading reply:", err)
		return
	}
	fmt.Printf("server says: %s", reply)
}
