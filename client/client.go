package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"
)

func ClientRun() {

	var choice string

	fmt.Print("tsl or not [y/N]: ")
	fmt.Scanln(&choice)

	conn := getConn(choice == "y")
	defer conn.Close()

	go func() {
		for {
			var msg string
			fmt.Scanln(&msg)
			if msg == "exit" {
				break
			}
			_, err := io.WriteString(conn, msg)
			if err != nil {
				log.Fatalf("client: write: %s", err)
				break
			}
		}
	}()

	for {
		d := 512
		buf := make([]byte, d)
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalf("client: read: %s", err)
			break
		}

		fmt.Printf("client: read %q\n", string(buf[:n]))
	}

}

func getConn(withTls bool) net.Conn {
	service := "0.0.0.0:8000"

	if withTls {
		cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
		if err != nil {
			log.Fatalf("server: loadkeys: %s", err)
		}
		config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
		conn, err := tls.Dial("tcp", service, &config)
		if err != nil {
			log.Fatalf("client: dial: %s", err)
		}
		fmt.Println("client: connected to: ", conn.RemoteAddr(), "with tls")

		state := conn.ConnectionState()
		for _, v := range state.PeerCertificates {
			fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
			fmt.Println(v.Subject)
		}
		log.Println("client: handshake: ", state.HandshakeComplete)
		log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

		return conn
	} else {
		conn, err := net.Dial("tcp", service)
		if err != nil {
			log.Fatalf("client: dial: %s", err)
		}
		fmt.Println("client: connected to: ", conn.RemoteAddr())
		return conn
	}

}
