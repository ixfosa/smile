package main

import (
	"fmt"
	"net"
)

/*******************UDP协议********************/
/*
UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联）
参考模型中一种无连接的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，
通常用于视频直播相关领
*/

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 19980,
	})
	if err != nil {
		fmt.Println("net.ListenUDP err :", err)
		return
	}

	defer listen.Close()

	for {
		var data [4096]byte
		n, addr, err := listen.ReadFromUDP(data[:])

		if err != nil {
			fmt.Println("net.ListenUDP err :", err)
			continue
		}

		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		listen.WriteToUDP([]byte("ixfosa"),addr)
	}
}



