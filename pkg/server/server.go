package server

import (
	c "Med/pkg/config"
	"bufio"
	"fmt"
	"log"
	"net"
)

func ServeAndReport(conf *c.Server) {
	address := conf.Host + ":" + conf.Port
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln("ERROR - SERVER:", err, "Address used:", address)
	}
	fmt.Println("Listening on", address)

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("ERROR - SERVER:", err, "(Accepting)")
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
		conn.Write([]byte("OK\n"))
	}
}
