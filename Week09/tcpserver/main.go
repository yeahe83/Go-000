package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	done := make(chan struct{})

	// server
	go server(done)

	// client
	go client()

	<-done

}

func server(done chan struct{}) {

	Listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tcp server listening")

	message := make(chan []byte)

	go func() {
		for {
			conn, err := Listener.Accept()
			defer conn.Close()
			if err != nil {
				log.Printf("%v\n", err)
				continue
			}

			writer := bufio.NewWriter(conn)
			reader := bufio.NewReader(conn)

			go func() {
				for {
					line, _, err := reader.ReadLine()
					if err != nil {
						log.Printf("%v\n", err)
					}
					if strings.Compare(string(line), "quit") == 0 {
						log.Println("bye!")
						done <- struct{}{}
					}
					message <- line
				}
			}()

			go func() {
				for {
					writer.WriteString("hello ")
					writer.Write(<-message)
					writer.WriteString("\n")
					writer.Flush()
				}
			}()
		}
	}()
}

func client() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			_, err = io.Copy(os.Stdout, conn)
			if err != nil {
				log.Printf("%v\n", err)
				continue
			}
		}
	}()

	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		log.Printf("%v\n", err)
	}
}

/*

go run main.go
2021/01/28 01:52:53 tcp server listening
111
hello 111
222
hello 222
333
hello 333
quit
2021/01/28 01:53:00 bye!

*/
