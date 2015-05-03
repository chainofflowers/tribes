package peers

// worked in playground, to make it safe

import (
	"../tools/"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"time"
)

var (
	client_keyfile string = "/News/tls/client.key"
	client_pemfile string = "/News/tls/client.crt"
	server_keyfile string = "/News/tls/server.key"
	server_pemfile string = "/News/tls/server.crt"
	TLS_engine     string = "Tribes TLS Engine 1.0"
)

func init() {

	var user_home = tools.GetHomeDir()
	client_keyfile = filepath.Join(user_home, client_keyfile)
	client_pemfile = filepath.Join(user_home, client_pemfile)
	server_keyfile = filepath.Join(user_home, server_keyfile)
	server_pemfile = filepath.Join(user_home, server_pemfile)
	log.Println("[TLS] Engine INIT: %s", TLS_engine)
}

func RotateKeysAndCert() {

	log.Println("[TLS] Key and Cert Renewal engine started")

	for {

		time.Sleep(1 * time.Hour)
		CreateKeysAndCert(tools.RandSeq(6), tools.RandSeq(8), tools.RandSeq(7))

	}

}

func CreateKeysAndCert(country string, organization string, unit string) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Country:            []string{country},
			Organization:       []string{organization},
			OrganizationalUnit: []string{unit},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		SubjectKeyId:          []byte{1, 2, 3, 4, 5},
		BasicConstraintsValid: true,
		IsCA:        true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	priv, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Println("[TLS] Create client Key failed: ", err)
		return
	}

	pub := &priv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("[TLS] Create client Certificate failed: ", err)
		return
	}
	ca_f := client_pemfile
	log.Println("[TLS] Writing client pemfile to: ", ca_f)
	ioutil.WriteFile(ca_f, ca_b, 0777)

	priv_f := client_keyfile
	priv_b := x509.MarshalPKCS1PrivateKey(priv)
	log.Println("[TLS] Writing client keyfile to: ", priv_f)
	ioutil.WriteFile(priv_f, priv_b, 0777)

	server_cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Country:            []string{country},
			Organization:       []string{organization},
			OrganizationalUnit: []string{unit},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	priv2, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Println("[TLS] Create server Key failed: ", err)
		return
	}

	pub2 := &priv2.PublicKey
	cert2_b, err2 := x509.CreateCertificate(rand.Reader, server_cert, ca, pub2, priv)
	if err2 != nil {
		log.Println("[TLS] Create server certificate failed: ", err2)
		return
	}

	cert2_f := server_pemfile
	log.Println("[TLS] Writing server certificate to: ", cert2_f)
	ioutil.WriteFile(cert2_f, cert2_b, 0777)

	priv2_f := server_keyfile
	priv2_b := x509.MarshalPKCS1PrivateKey(priv2)
	log.Println("[TLS] Writing server key to: ", priv2_f)
	ioutil.WriteFile(priv2_f, priv2_b, 0777)

	ca_c, _ := x509.ParseCertificate(ca_b)
	cert2_c, _ := x509.ParseCertificate(cert2_b)

	err3 := cert2_c.CheckSignatureFrom(ca_c)
	log.Println("[TLS] Signature check returned: ", err3 == nil)
}
