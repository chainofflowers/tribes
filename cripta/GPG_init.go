package cripta

import (
	"log"
	"tribes/config"
	"tribes/tools"
)

func init() {

	var thekey string
	var metadata map[string]string
	metadata = make(map[string]string)
	log.Println("[GPG] Engine started")

	if thekey = config.GetTribeID(); len(thekey) < 32 {
		log.Println("[GPG] EEK: TribeID %d ", len(thekey))
		log.Println("[GPG] EEK: TribeID shorter than 32 bytes. Generating a random one")
		thekey = tools.RandSeq(42) // 42 because of yes.
		log.Println("[GPG] your 1-node tribe is: " + thekey)
	}

	log.Println("[GPG] TribeID is: " + thekey)
	test_cleartext := tools.RandSeq(128)
	metadata["Command"] = "NOOP"
	metadata["Group"] = "it.doesnt.exists"
	log.Println("[GPG] GPG Integrity test initiated")
	log.Println("[GPG] GPG Encrypted payload below: \n", GpgEncrypt(test_cleartext, metadata))
	log.Println("[GPG] GPG Integrity Test passed: ", test_cleartext == GpgDecrypt(GpgEncrypt(test_cleartext, metadata)))
	d_Command := GpgGetHeaders(GpgEncrypt(test_cleartext, metadata))
	log.Println("[GPG] GPG Serialization Test #1 passed: ", d_Command["Command"] == metadata["Command"])
	log.Println("[GPG] GPG Serialization Test #2 passed: ", d_Command["Group"] == metadata["Group"])
}

func GPG_Engine_Start() {
	log.Println("[GPG] Invoked")

}
