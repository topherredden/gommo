package main

import (
	"io"
	"log"
	"net"
    "./protocol"
    //"encoding/binary"
)

var _ = io.Copy

func main() {
	l, err := net.Listen("tcp", ":2593")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
            continue
		}

        go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
    buf := make([]byte, 4096)

    for {
        n, err := c.Read(buf)
        if err != nil || n == 0 {
            c.Close()
            break
        }

        n, err = c.Write(buf[0:n])
        if err != nil {
            c.Close()
            break
        }

        if n <= 0 {
            continue
        }

        // Read command
        cmd := buf[0]

        if cmd == protocol.Packet_AccountLogin {
            //c.Write([]byte(0x82))
            response := make([]byte, 2)
            response[0] = 0x82
            response[1] = 0x01


            log.Printf("Packet: %#x", response[0])

            //binary.Write(c, binary.BigEndian, response)
            w, err := c.Write(response)

            log.Println(w)

            if err != nil {
                log.Println(err)
                c.Close()
                break
            }
        }

        log.Printf("Packet: %#x", cmd)
        //log.Println(buf[0:n])
    }

    log.Printf("Connection from %v closed.", c.RemoteAddr())
}
