package peers

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

var (
	runserver   = make(chan bool) // runserver <- false to stop the server
	server_dead sync.WaitGroup    // server_dead.Wait() to be sure the server dies.

)

// worked in playground, now  make it safe

func TLS_Frontend(local_ipandport string) {

	ca_b, err := ioutil.ReadFile(server_pemfile)

	if err != nil {
		log.Println("[TLS] Cannot read server certificate %s : %s", server_pemfile, err)
		return
	}

	ca, err := x509.ParseCertificate(ca_b)
	if err != nil {
		log.Println("[TLS] Cannot parse server certificate %s: %s", server_pemfile, err)
		return
	}

	priv_b, err := ioutil.ReadFile(server_keyfile)
	if err != nil {
		log.Println("[TLS] Cannot read server key %s : %s", server_keyfile, err)
		return
	}

	priv, err := x509.ParsePKCS1PrivateKey(priv_b)
	if err != nil {
		log.Println("[TLS] Cannot parse server key %s : %s", server_keyfile, err)
		return
	}

	pool := x509.NewCertPool()
	pool.AddCert(ca)

	cert := tls.Certificate{
		Certificate: [][]byte{ca_b},
		PrivateKey:  priv,
	}

	config := tls.Config{
		ClientAuth:   tls.RequireAnyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    pool,
	}
	config.Rand = rand.Reader

	// to put my IP and port of the tribe

	listener, err := tls.Listen("tcp", local_ipandport, &config)
	if err != nil {
		log.Println("[TLS] Server cannot listen ar %s: %s", local_ipandport, err)
	} else {
		log.Println("[TLS] Server listening at %s: ", local_ipandport)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[TLS] Server cannot answer: %s", err)
			break
		} else {
			defer conn.Close()
			log.Println("[TLS] Server  accepted client from %s", conn.RemoteAddr())
			go handleClient(conn)
		}
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

		tlscon, ok := conn.(*tls.Conn)
		if ok {
			state := tlscon.ConnectionState()
			sub := state.PeerCertificates[0].Subject
			log.Println(sub)
		}

		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		n, err = conn.Write(buf[:n])

		n, err = conn.Write(buf[:n])
		log.Printf("server: conn: wrote %d bytes", n)

		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")
}
