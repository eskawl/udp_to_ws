# udp_to_ws
udp_to_ws provides udp data over websocket

## Usage

### Connect to a UDP source

	conn := ConnectUDP("localhost:20777")

### Implement the handler interface

The UdpWebSocket method needs to be implemented

	type socketHandlers struct {
		conn *net.UDPConn
	}

	func (h *socketHandlers) UdpWebSocket(c *websocket.Conn, w http.ResponseWriter, r *http.Request) {
		packetData := make([]byte, 4*10) // 10 bytes
		extraData := make([]byte, 5*10) // 5 bytes
		ReadUDP(h.conn, packetData, extraData)

		c.WriteMessage(websocket.BinaryMessage, packetData)
	}

### Start the server

	handlers := &socketHandlers{
		conn: conn,
	}
	StartServer(handlers)