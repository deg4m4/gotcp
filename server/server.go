package server

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"
)

func ServerRun() {

	var choice string

	fmt.Print("tsl or not [y/N]: ")
	fmt.Scanln(&choice)

	listener := getListner(choice == "y")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
		}
		go handleClient(conn)
	}

}

func getListner(withTls bool) net.Listener {
	service := "0.0.0.0:8000"

	if withTls {
		cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
		if err != nil {
			panic(err)
		}

		config := tls.Config{Certificates: []tls.Certificate{cert}}
		config.Rand = rand.Reader

		listener, err := tls.Listen("tcp", service, &config)
		if err != nil {
			log.Fatalf("server: listen: %s", err)
		}

		log.Print("server: listening with lts")

		return listener
	} else {
		listner, err := net.Listen("tcp", service)
		if err != nil {
			log.Fatalf("server: listen: %s", err)
		}

		log.Print("server: listening")

		return listner
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("server: conn: read: %s", err)
			break
		}
		log.Printf("server: conn: echo %q\n", string(buf[:n]))

		if string(buf[:n]) == "gh" {

			n, err = io.WriteString(conn, "You call special keyword ✨\n")
			log.Printf("server: conn: wrote %d bytes", n)

			if err != nil {
				log.Printf("server: write: %s", err)
				break
			}
		}
	}
	log.Println("server: conn: closed")

}
