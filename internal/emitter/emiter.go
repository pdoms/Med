package emitter

import (
	c "Med/pkg/config"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

const (
	SOB = 11
	EOB = 28
	CR  = 13
)

type EmitArgs struct {
	Protocol   string
	SendSample bool
}

func ProcessArgs() EmitArgs {
	var args EmitArgs
	prot := flag.String("protocol", "mllp", "defines transfer protocol")
	sample := flag.Bool("test", false, "emits one random test message")
	flag.Parse()
	args.Protocol = *prot
	args.SendSample = *sample
	return args

}

func GetTestMsg() (msg string, err error) {
	file, err := os.Open("/home/paulo/Projects/Med/internal/emitter/testmsg.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	return string(bytes), err
}

func Mllp(msg string) []byte {
	l := len(msg)
	msgBytes := make([]byte, 3+l)
	msgBytes[0] = SOB
	copy(msgBytes[1:], []byte(msg))
	end := []byte{EOB, CR}
	copy(msgBytes[l+1:], end)
	return msgBytes
}

func Emit(conf *c.Emitter) {
	args := ProcessArgs()
	var toSend []byte
	if conf.Protocol == "" {
		conf.Protocol = args.Protocol
		fmt.Println("INFO - EMITTER: using", conf.Protocol)
	}

	if args.SendSample {
		msg, err := GetTestMsg()
		if err != nil {
			fmt.Println("Error - EMITTER", err)
		}
		toSend = Mllp(msg)
	}
	servAddress := fmt.Sprintf("%s:%s", conf.Connection.Host, conf.Connection.Port)
	tcpAddress, err := net.ResolveTCPAddr("tcp", servAddress)
	if err != nil {
		log.Fatalln("ERROR - EMITTER: Unable to resolve address,", tcpAddress, err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		log.Fatalln("ERROR - EMITTER: Unable to dial,", err)
	}
	defer conn.Close()
	responseReader := bufio.NewReader(conn)

	for {
		if _, err = conn.Write(toSend); err != nil {
			log.Fatalln("ERROR - EMITTER: Unable to write request,", err)
		}
		response, err := responseReader.ReadString('\n')
		if err != nil {
			log.Fatalln("ERROR - EMITTER: crooked response,", err)
		}
		fmt.Println("RESPONSE", strings.TrimSpace(response))
		os.Exit(0)
	}
}
