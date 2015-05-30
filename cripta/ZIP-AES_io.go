package cripta

type JsonPack struct {
	Clear_text string
	Crypted    string
	aes_key    string
}

// quick and easy interface to send and receive encrypted+zipped JSON
// just declare a new JsonPack and use it.

func (this *JsonPack) ZipAndCrypt() {

	this.Crypted = EasyCrypt(TextZip(this.Clear_text), this.aes_key)

}

func (this *JsonPack) DeCryptAndUnzip() {

	this.Clear_text = TextUnzip(EasyDeCrypt(this.Crypted, this.aes_key))

}
