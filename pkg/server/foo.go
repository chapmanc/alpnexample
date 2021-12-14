package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func GetFooProtocol(shutdownCh chan struct{}) func(server *http.Server, conn *tls.Conn, handler http.Handler) {
	return func(server *http.Server, conn *tls.Conn, handler http.Handler) {

		//fmt.Printf("handle foo started, server:\n%#v\nconn:\n%#v\nhandler:\n%#v\n", *server, *conn, handler)
		fmt.Printf("handle foo started\n")

		rdr := bufio.NewReader(conn)
		wtr := bufio.NewWriter(conn)
		for {
			select {
			case <-shutdownCh:
				return
			default:
				conn.SetReadDeadline(time.Now().Add(2 * time.Second))
				line, err := rdr.ReadString('\n')
				if err != nil {
					fmt.Printf("error during foo server reading: %v\n", err)
					continue
				}

				fmt.Printf("foo server got %s\n", line)

				conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
				_, err = wtr.WriteString(line)
				if err != nil {
					fmt.Printf("error during foo client writing: %v\n", err)
				}
				err = wtr.Flush()
				if err != nil {
					fmt.Printf("error during foo client write flush: %v\n", err)
					continue
				}
				fmt.Printf("foo server wrote %s\n", line)
			}
		}
	}
}

// func GetFoo(wg *sync.WaitGroup, shutdownCh chan struct) (fn func(server *http.Server, conn *tls.Conn, handler http.Handler), error) {

// 	return  func(server *http.Server, conn *tls.Conn, handler http.Handler) {

// 	}

// }
