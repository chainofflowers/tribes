package tribe

import (
	"../config/"
	"../tools/"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
)

type Push_Message struct {
	Command   string `json:"command"`
	MessageID string `json:"messageid"`
	Headers   string `json:"headers-b64"`
	Body      string `json:"body-b64"`
	Xover     string `json:"xover-b64"`
	Fill      string `json:"fill"` // fill with 32 random chars, to match the AES pad
}

//var Tbe_Pack Push_Message
//Tbe_Pack.Command = "POST"
//Tbe_Pack.MessageID ="TUAMADREID"
//Tbe_Pack.Headers = "NegroBianco"  // base64 encoded
//Tbe_Pack.Body = TUAMADRETUAMADRETUAMADRETUAMADRE  // base64 encoded

func (this *Push_Message) Post2AES_JSON(messageid string, body string, headers string, xover string) string {

	var EncPost MyEncryption

	// some limit on body, headers and xover size

	body = tools.ShortenString(body, 50000)
	headers = tools.ShortenString(body, 10000)
	xover = tools.ShortenString(xover, 5000)

	messageid_b64 := base64.StdEncoding.EncodeToString([]byte(messageid))
	body_b64 := base64.StdEncoding.EncodeToString([]byte(body))
	headers_b64 := base64.StdEncoding.EncodeToString([]byte(headers))
	xover_b64 := base64.StdEncoding.EncodeToString([]byte(xover))

	// make everything a ASCII string

	this.Command = "POST"
	this.Body = body_b64
	this.Headers = headers_b64
	this.MessageID = messageid_b64
	this.Xover = xover_b64
	this.Fill = tools.RandSeq(32)

	// serialize in JSON

	Packed_MSG, _ := json.Marshal(this)

	// made it AES Armored

	copy(EncPost.MyText_cleartext, []byte(Packed_MSG))
	copy(EncPost.MyAES_key, []byte(config.GetTribeID()))
	EncPost.AESencrypt()
	return this.Command + "\n" + EncPost.AESArmored()

}

// decrypts a "POST" body

func (this *Push_Message) AES_JSON2Post(jbod string) {

	var DecPost MyEncryption

	copy(DecPost.MyAES_key, []byte(config.GetTribeID()))
	bin_post := base64.StdEncoding.DecodeString(jbod)
	copy(DecPost.MyText_encrypted, []byte(bin_post))
	DecPost.AESdecrypt()

	err := json.Unmarshal(DecPost.MyText_cleartext, &this)

	if err != nil {
		this.Command = "NOOP"
	}

}

// writes on UDP socket the encoded message

func SendPost(conn net.UDPConn, messageid string) {

	// reads h,b and x files , creates the block and sends to the peer

	// TODO

}

func ReceivePost(conn net.UDPConn, jbod string) {

	// takes the jbod , decypts and writes the x,h,b files

	//TODO

}
