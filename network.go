package udp_to_ws

import (
	"log"
	"net"
)

func ConnectUDP(port string) *net.UDPConn {
	udpAddress, err := net.ResolveUDPAddr("udp4", port)

	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp4", udpAddress)

	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func ReadUDP(conn *net.UDPConn, packetData []byte, extraData []byte) {
	_, _, _, _, err := conn.ReadMsgUDP(packetData, extraData)

	if err != nil {
		log.Println("UDP Read error: ", err)
	}
}
