package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/********************UDP客户端*******************/
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 19980,
	})

	if err != nil {
		fmt.Println("net.DialUDP err: ", err)
	}

	defer socket.Close()
	for  {
		input := bufio.NewReader(os.Stdin)
		rendData, _:= input.ReadString('\n')

		_, err := socket.Write([]byte(rendData))

		if err != nil {
			fmt.Println("bufio.NewReader err: ", err)
			continue
		}
		buf := [4096]byte{}
		n, addr, err := socket.ReadFromUDP(buf[:])

		if err != nil {
			fmt.Println("socket.ReadFromUDP err: ", err)
			continue
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(buf[:n]), addr, n)
	}


}


