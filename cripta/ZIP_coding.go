package cripta

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"
)

func TextZip(text string) string {

	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	zlib.NewWriterLevel(w, 9)
	w.Write([]byte(text))
	w.Close() // You must close this first to flush the bytes to the buffer.
	return b.String()

}

func TextUnzip(text string) string {

	deflated := bytes.NewReader([]byte(text))

	enflated, err := zlib.NewReader(deflated)

	if err != nil {
		log.Println("[ZIP] Input was not compressed")
		return text
	}

	if s, err := ioutil.ReadAll(enflated); err == nil {
		return string(s)
	} else {
		log.Println("[ZIP] Empty content ")
		return ""
	}

	return text

}
