package cripta

import (
	"../config/"
	"../tools/"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"os"
)

type MyEncryption struct {
	MyAES_key        []byte
	MyAES_Error      error
	MyText_encrypted []byte
	MyText_cleartext []byte
	MyCypherBlock    cipher.Block
	MyIv             []byte
}

func init() {

	var LocalTribe MyEncryption
	var thekey string
	log.Println("[AES] Engine started")

	if thekey = config.GetTribeID(); len(thekey) != 32 {
		log.Println("[AES] EEK: TribeID %d ", len(thekey))
		log.Println("[AES] EEK: TribeID shorter than 32 bytes. Generating a random one")
		thekey = tools.RandSeq(32)
		log.Println("[AES] your 1-node tribe is: " + thekey)
	}

	LocalTribe.MyAES_key = []byte(thekey)
	log.Println("[AES] TribeID is: ", string(LocalTribe.MyAES_key))
	LocalTribe.MyText_cleartext = []byte(tools.RandSeq(33))
	log.Println("[AES] TribeGreeting before is: " + string(LocalTribe.MyText_cleartext))
	LocalTribe.AESencrypt()
	log.Println("[AES] test Encryption executed")
	copy(LocalTribe.MyText_cleartext, []byte("WRONG->"))
	LocalTribe.AESdecrypt()
	log.Println("[AES] TribeGreeting after  is: " + string(LocalTribe.MyText_cleartext))
	log.Println("[UDP] Now Starting UDP listener ")

}

func AES_Engine_Start() {
	log.Println("[AES] Invoked")

}

func (this *MyEncryption) AESencrypt() {
	this.MyCypherBlock, this.MyAES_Error = aes.NewCipher(this.MyAES_key)
	if this.MyAES_Error != nil {
		log.Println("[AES] Cannot add NewCipher: ", this.MyAES_Error)
		return
	}
	b := base64.StdEncoding.EncodeToString(this.MyText_cleartext)

	this.MyText_encrypted = make([]byte, aes.BlockSize+len(b))
	this.MyIv = this.MyText_encrypted[:aes.BlockSize]

	if _, this.MyAES_Error = io.ReadFull(rand.Reader, this.MyIv); this.MyAES_Error != nil {
		log.Println("[AES] Cannot create iv from rand: ", this.MyAES_Error)
		return
	}
	cfb := cipher.NewCFBEncrypter(this.MyCypherBlock, this.MyIv)
	cfb.XORKeyStream(this.MyText_encrypted[aes.BlockSize:], []byte(b))
	return
}

//=================
// Decryption function
//=================

func (this *MyEncryption) AESdecrypt() {
	this.MyCypherBlock, this.MyAES_Error = aes.NewCipher(this.MyAES_key)
	if this.MyAES_Error != nil {
		log.Println("[AES] Cannot add NewCipher: ", this.MyAES_Error)
		return
	}

	if len(this.MyText_encrypted) < aes.BlockSize {
		this.MyAES_Error = errors.New("Ciphertext too short")
		log.Println("[AES] Ciphertext too short")
		return
	}
	this.MyIv = this.MyText_encrypted[:aes.BlockSize]
	this.MyText_encrypted = this.MyText_encrypted[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(this.MyCypherBlock, this.MyIv)
	cfb.XORKeyStream(this.MyText_encrypted, this.MyText_encrypted)
	this.MyText_cleartext, this.MyAES_Error = base64.StdEncoding.DecodeString(string(this.MyText_encrypted))
	if this.MyAES_Error != nil {
		log.Println("[AES] Error while decrypting ", this.MyAES_Error)
		this.MyText_cleartext = nil

	}
	return
}

//
// Returns armored AES encrypted
//

func (this *MyEncryption) AESArmored() string {

	return base64.StdEncoding.EncodeToString(this.MyText_encrypted)

}

func EasyCrypt(text, key string) string {

	var aes_tmp MyEncryption

	aes_tmp.MyText_cleartext = []byte(text)
	aes_tmp.MyAES_key = []byte(key)
	aes_tmp.AESencrypt()
	return aes_tmp.AESArmored()

}

func EasyDeCrypt(text, key string) string {

	var aes_tmp MyEncryption
	aes_tmp.MyText_encrypted = []byte(text)
	aes_tmp.MyAES_key = []byte(key)
	aes_tmp.AESdecrypt()
	return string(aes_tmp.MyText_cleartext)

}