package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, _ := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: ")
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) *Client {
	b := bufio.NewReader(client)
	w := bufio.NewWriter(client)
	return NewClient(b, w, client)

}
func main() {
	var new_room = *NewRoom()
	server, _ := net.Listen("tcp", ":"+strconv.Itoa(5000))
	if server == nil {
		panic("couldn't start listening: ")
	}
	conns := clientConns(server)
	go new_room.run()
	for {
		new_client := handleConn(<-conns)
		new_room.AddUser(new_client)
	}
}
