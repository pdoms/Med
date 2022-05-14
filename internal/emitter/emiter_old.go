package emitter

/* import (
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

func EmitOLD(conf *c.Emitter) {
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

		servAddress := fmt.Sprintf("%s:%s", conf.Connection.Host, conf.Connection.Port)
		tcpAddress, err := net.ResolveTCPAddr("tcp", servAddress)
		if err != nil {
			log.Fatalln("ERROR - EMITTER: Unable to resolve address,", tcpAddress, err)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddress)
		if err != nil {
			log.Fatalln("ERROR - EMITTER: Unable to dial,", err)
		}
		responseReader := bufio.NewReader(conn)
		defer conn.Close()
		for {
			if _, err = conn.Write(toSend); err != nil {
				log.Fatalln("ERROR - EMITTER: Unable to write request,", err)
			}

			response, err := responseReader.ReadString('\n')
			if err != nil {
				log.Fatalln("ERROR - EMITTER: crooked response,", err)
			}
			fmt.Println("RESPONSE", strings.TrimSpace(response))
			if _, err = conn.Write(toSend); err != nil {
				log.Fatalln("ERROR - EMITTER: Unable to write request,", err)
			}
			os.Exit(0)
		}
	} else {
		servAddress := fmt.Sprintf("%s:%s", conf.Connection.Host, conf.Connection.Port)
		tcpAddress, err := net.ResolveTCPAddr("tcp", servAddress)
		if err != nil {
			log.Fatalln("ERROR - EMITTER: Unable to resolve address,", tcpAddress, err)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddress)
		if err != nil {
			log.Fatalln("ERROR - EMITTER: Unable to dial,", err)
		}

		for {
			fmt.Println("Enter Message:")
			var lines []string
			for scn.Scan() {
				line := scn.Text()
				if len(line) == 1 {
					if line[0] == ' ' {
						break
					}
				}
				lines = append(lines, line)
			}

			if len(lines) > 0 {
				fmt.Println()
				fmt.Println("Processing Message to be send...")
				msg := []byte{SOB}
				for _, line := range lines {
					lineBytes := []byte(line)
					msg = append(msg, lineBytes...)
				}
				msg = append(msg, EOB)
				msg = append(msg, CR)
				responseReader := bufio.NewReader(conn)
				defer conn.Close()
				for {
					if _, err = conn.Write(msg); err != nil {
						log.Fatalln("ERROR - EMITTER: Unable to write request,", err)
					}
					fmt.Println("WHAT", responseReader)
					response, err := responseReader.ReadString('\n')
					fmt.Println(response)
					if err != nil {
						log.Fatalln("ERROR - EMITTER: crooked response,", err)
					}
					fmt.Println("RESPONSE", strings.TrimSpace(response))
					break
				}
				fmt.Println("HERE")
			}
			if err := scn.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

	}

}
*/
