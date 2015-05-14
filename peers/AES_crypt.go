package peers

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
	log.Println("[AES] Engine started")
	thekey := config.GetTribeID()

	if len(thekey) != 32 {
		log.Println("[AES] EEK: TribeID shorter than 32 bytes. Cannot start tribes")
		os.Exit(3)
	}

	LocalTribe.MyAES_key = []byte(config.GetTribeID())
	log.Println("[AES] TribeID is: ", string(LocalTribe.MyAES_key))
	LocalTribe.MyText_cleartext = []byte(tools.RandSeq(33))
	log.Println("[AES] TribeGreeting before is: " + string(LocalTribe.MyText_cleartext))
	LocalTribe.encrypt()
	log.Println("[AES] test Encryption executed")
	LocalTribe.decrypt()
	log.Println("[AES] TribeGreeting after  is: " + string(LocalTribe.MyText_cleartext))
}

func AES_Engine_Start() {
	log.Println("[AES] Invoked AES Engine")

}

func (this *MyEncryption) encrypt() {
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

func (this *MyEncryption) decrypt() {
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
