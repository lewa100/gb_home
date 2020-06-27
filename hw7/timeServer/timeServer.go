package timeServer

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strings"
	"time"
)

// TimeServer - Show time for client.
func TimeServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

// checkCMD - Connect is close if input "exit".
func checkCMD(c net.Conn) {
	buff := make([]byte, 1024)
	cmd := "exit"

	c.Read(buff)
	str := strings.ToLower(strings.Replace(string(buff), "\r\n", "", -1))
	//fmt.Println(str)
	if strings.Contains(str, cmd) {
		c.Close()
	}
}

// handleConn - working with connect.
func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		checkCMD(c)
		time.Sleep(1 * time.Second)
	}
}
