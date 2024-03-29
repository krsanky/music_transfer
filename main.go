package main

import (
	"fmt"
	"os"
	"time"

	"nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/rep"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func usage() {
	fmt.Printf("usage....\n:")
}

func main() {
	//for i, a := range os.Args[1:] {
	//	fmt.Printf("%d:%s ", i, a)
	//}

	if len(os.Args) >= 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "web":
			//web()
		case "db":
			//dbstuff()
		case "date":
			fmt.Printf("date:%s\n", date())
		default:
			usage()
		}
	} else {
		usage()
	}
}

//url=tcp://127.0.0.1:40899
//./reqrep node0 $url & node0=$! && sleep 1
//./reqrep node1 $url
//kill $node0
func node0(url string) {
	var sock mangos.Socket
	var err error
	var msg []byte
	if sock, err = rep.NewSocket(); err != nil {
		die("can't get new rep socket: %s", err)
	}
	if err = sock.Listen(url); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}
	for {
		// Could also use sock.RecvMsg to get header
		msg, err = sock.Recv()
		if string(msg) == "DATE" { // no need to terminate
			fmt.Println("NODE0: RECEIVED DATE REQUEST")
			d := date()
			fmt.Printf("NODE0: SENDING DATE %s\n", d)
			err = sock.Send([]byte(d))
			if err != nil {
				die("can't send reply: %s", err.Error())
			}
		}
	}
}
