package server

import (
	c "Med/pkg/config"
	"Med/pkg/hl7"
	"fmt"
	"log"
	"net"
)

func ServeAndReportMllp(conf *c.Server) {
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
			return
		}
		go handleMllpRequest(conn)

	}

}

func handleMllpRequest(conn net.Conn) {
	fmt.Println("INFO - SERVER: received mllp message from", conn.RemoteAddr())
	defer conn.Close()
	scanner := hl7.NewMllpScanner(conn)
	err := scanner.Scan()
	if err != nil {
		//write response
		return
	}
	//logic to do something with the message

	fmt.Println(string(scanner.Msg))
	conn.Write([]byte("OK\n"))
	fmt.Println("INFO - SERVER: handled request - ready for more")
	//buf := make([]byte, 12)
	//for {
	//	len, err := conn.Read(buf)
	//	if err != nil {
	//		fmt.Println("ERROR - SERVER: ", err)
	//	}
	//	fmt.Println(buf[:len])
	//}

}
