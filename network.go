package udp_to_ws

import (
	"log"
	"net"
)

// Creates a UDP Connection to the given port. The port is passed
// to net.ResolveUDPAddr to start the connection
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

// Reads the data available on the UDP connection
// Data will be read into packet data and extra data
func ReadUDP(conn *net.UDPConn, packetData []byte, extraData []byte) {
	_, _, _, _, err := conn.ReadMsgUDP(packetData, extraData)

	if err != nil {
		log.Println("UDP Read error: ", err)
	}
}
