package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	util "github.com/meetthehorizon/scratchpad/tcp-echo-server/internal"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("connected to client: %v\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	req, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}

	rep := strings.ToUpper(req)
	if _, err := conn.Write([]byte(rep + "\n")); err != nil {
		fmt.Print(err)
	}
}

func main() {
	serverAddr := util.GetServerAddr()
	server, err := net.Listen("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on %v...\n", serverAddr)

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("encountered error when accepting connction", err)
		}

		go handleConnection(conn)
	}
}
