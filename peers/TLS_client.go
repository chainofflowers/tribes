package peers

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

// worked with the playground . to adapt

func TLS_socket_Create(peer_addressandport string) tls.Conn {
	cert2_b, err := ioutil.ReadFile(client_pemfile)
	if err != nil {
		log.Println("[TLS] Cannot read client certificate %s : %s", client_pemfile, err)
	}

	priv2_b, err := ioutil.ReadFile(client_keyfile)
	if err != nil {
		log.Println("[TLS] Cannot read client key %s : %s", client_keyfile, err)

	}

	priv2, err := x509.ParsePKCS1PrivateKey(priv2_b)
	if err != nil {
		log.Println("[TLS] Cannot parse client key : %s", err)

	}

	cert := tls.Certificate{
		Certificate: [][]byte{cert2_b},
		PrivateKey:  priv2,
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", peer_addressandport, &config)
	if err != nil {
		log.Println("[TLS] Client problem on dial: %s", err)

	}

	log.Println("[TLS] client: connected to ", conn.RemoteAddr())

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {

		log.Println("[TLS] Subject: ", v.Subject)
	}
	log.Println("[TLS] Client: handshake: ", state.HandshakeComplete)
	log.Println("[TLS] Client: mutual: ", state.NegotiatedProtocolIsMutual)

	return *conn

}
