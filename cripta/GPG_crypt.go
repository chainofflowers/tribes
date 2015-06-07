package cripta

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"tribes/config"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func GpgEncrypt(text2encrypt string, headers map[string]string) string {
	var encrypted bytes.Buffer
	foo := bufio.NewWriter(&encrypted)
	w, _ := armor.Encode(foo, "TRIBES PAYLOAD", headers)
	plaintext, _ := openpgp.SymmetricallyEncrypt(w, []byte(config.GetTribeID()), nil, nil)
	fmt.Fprintf(plaintext, text2encrypt)
	plaintext.Close()
	w.Close()
	foo.Flush()
	return encrypted.String()
}

func GpgGetHeaders(text2decrypt string) map[string]string {

	encrypted := bytes.NewBuffer([]byte(text2decrypt))
	var tmp_map map[string]string
	tmp_map = make(map[string]string)

	bin_encrypt, err := armor.Decode(encrypted)
	if err != nil {
		log.Printf("[PGP] not an armored payload: %s", err.Error())
		tmp_map["Command"] = "NOOP"
		return tmp_map
	}

	return bin_encrypt.Header

}

func GpgDecrypt(text2decrypt string) string {
	encrypted := bytes.NewBuffer([]byte(text2decrypt))

	bin_encrypt, err := armor.Decode(encrypted)
	if err != nil {
		log.Printf("[PGP] not an armored payload: %s", err.Error())
		return ""
	}

	cleartext_md, err := openpgp.ReadMessage(bin_encrypt.Body, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		return []byte(config.GetTribeID()), nil
	}, nil)
	if err != nil {
		log.Printf("[PGP] Can't decrypt payload: %s", err.Error())
		return ""
	}

	plaintext, err := ioutil.ReadAll(cleartext_md.UnverifiedBody)
	if err != nil {
		log.Printf("[PGP] Can't read cleartext: %s", err.Error())
		return ""
	}

	return string(plaintext)
}
