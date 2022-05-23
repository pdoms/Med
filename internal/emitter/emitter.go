package emitter

import (
	"Med/internal/protocols/mllp"
	u "Med/internal/utils"
	"Med/pkg/config"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type UseEmitter struct {
	msg  []byte
	mllp []byte
}

func (emit *UseEmitter) makeMllp() {
	l := len(emit.msg)
	msgBytes := make([]byte, l+3)
	msgBytes[0] = mllp.SOB
	copy(msgBytes[1:], emit.msg)
	copy(msgBytes[l+1:], []byte{mllp.EOB, mllp.CR})
	emit.mllp = msgBytes
}

func getTestMsg() (msg []byte, err error) {
	file, err := os.Open("/home/paulo/Projects/Med/internal/emitter/testmsg.txt")
	u.HandleError(err)
	msg, err = ioutil.ReadAll(file)
	defer file.Close()
	return msg, err
}

func NewEmitter() UseEmitter {
	var emitter UseEmitter
	isTest := flag.Bool("sngl", false, "sends one single hl7 message and then exits.")
	flag.Parse()
	var err error
	if *isTest {
		emitter.msg, err = getTestMsg()
		u.HandleError(err)
	} else {
		emitter.msg = make([]byte, 0)
	}

	return emitter
}

func promptMessage() {
	fmt.Println("Please enter message and then hit enter followed by Space:")
	fmt.Println()

}

func runAsRepl(emit *UseEmitter, conf *config.Emitter) {
	servAddress := fmt.Sprintf("%s:%s", conf.Connection.Host, conf.Connection.Port)
	tcpAddress, err := net.ResolveTCPAddr("tcp", servAddress)
	u.HandleError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	u.HandleError(err)
	defer conn.Close()
	responseReader := bufio.NewReader(conn)
	for {
		promptMessage()
		scn := bufio.NewScanner(os.Stdin)
	scanning:
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				if line[0] == ' ' {
					break scanning
				}
			}
			line += "\n"
			emit.msg = append(emit.msg, []byte(line)...)
		}
		fmt.Println()
		emit.makeMllp()
		if _, err = conn.Write(emit.mllp); err != nil {
			u.HandleError(err)
		}
		response, err := responseReader.ReadString('\n')
		u.HandleError(err)
		if strings.TrimSpace(response) == "OK" {
			fmt.Println("Message successfully sent.")
			fmt.Println()
		} else {
			fmt.Println(response)
			log.Fatal("ERROR:", err)
		}
	}
}

func sendSnglMsg(emit *UseEmitter, conf *config.Emitter) {
	emit.makeMllp()
	servAddress := fmt.Sprintf("%s:%s", conf.Connection.Host, conf.Connection.Port)
	tcpAddress, err := net.ResolveTCPAddr("tcp", servAddress)
	u.HandleError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	u.HandleError(err)
	defer conn.Close()
	responseReader := bufio.NewReader(conn)
	for {
		if _, err = conn.Write(emit.mllp); err != nil {
			log.Fatalln("ERROR:", err)
		}
		response, err := responseReader.ReadString('\n')
		u.HandleError(err)
		if strings.TrimSpace(response) == "OK" {
			fmt.Println("Message emitted.")
			os.Exit(0)
		} else {
			log.Fatal("ERROR:", err)
		}
	}
}

func Emit(conf *config.Emitter) {
	emit := NewEmitter()
	if len(emit.msg) == 0 {
		runAsRepl(&emit, conf)
	} else {
		sendSnglMsg(&emit, conf)
	}
}
