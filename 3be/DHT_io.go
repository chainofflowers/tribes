package tribe

import (
	"log"

	"github.com/secondbit/wendy"

	"time"
	dht "tribes/cripta"
)

func SliceToString(slice []string) string {

	var tmpStr string

	for i := range slice {

		tmpStr += slice[i] + "\n"

	}

	return tmpStr

}

func BCastMsgHeaders(payload []string, group string, messageID string) {

	var msgHdr map[string]string
	msgHdr = make(map[string]string)

	const layout = "0601021504"
	orario := time.Now()

	msgHdr[TRIBES_H_CMD] = TRIBES_HEADER
	msgHdr[TRIBES_H_GID] = group
	msgHdr[TRIBES_H_MID] = messageID + "@" + orario.Format(layout)

	msgPayload := SliceToString(payload) + "\r\n.\r\n"

	wendyMsgString := dht.GpgEncrypt(msgPayload, msgHdr)

	id, err := wendy.NodeIDFromBytes([]byte("FakeID"))
	if err != nil {
		log.Printf("[DHT-BCAST] Error creating fakenodeid %s", err.Error())
	}

	msg := cluster.NewMessage(byte(30), id, []byte(wendyMsgString))

	WendyBroadcast(msg)

	log.Printf("[DHT-BCAST] spreading around XOVER for %s", messageID)

}

func BCastMsgBody(payload []string, group string, messageID string) {

	var msgHdr map[string]string
	msgHdr = make(map[string]string)

	const layout = "0601021504"
	orario := time.Now()

	msgHdr[TRIBES_H_CMD] = TRIBES_BODY
	msgHdr[TRIBES_H_GID] = group
	msgHdr[TRIBES_H_MID] = messageID + "@" + orario.Format(layout)

	msgPayload := SliceToString(payload) // we have this already in the body.

	wendyMsgString := dht.GpgEncrypt(msgPayload, msgHdr)

	id, err := wendy.NodeIDFromBytes([]byte("FakeID"))
	if err != nil {
		log.Printf("[DHT-BCAST] Error creating fake nodeid %s", err.Error())
	}

	msg := cluster.NewMessage(byte(30), id, []byte(wendyMsgString))

	WendyBroadcast(msg)

	log.Printf("[DHT-BCAST] spreading around XOVER for %s", messageID)

}

func BCastMsgXover(payload []string, group string, messageID string) {

	var msgHdr map[string]string
	msgHdr = make(map[string]string)

	const layout = "0601021504"
	orario := time.Now()

	msgHdr[TRIBES_H_CMD] = TRIBES_XOVER
	msgHdr[TRIBES_H_GID] = group
	msgHdr[TRIBES_H_MID] = messageID + "@" + orario.Format(layout)

	msgPayload := SliceToString(payload) // we are printing it already when requested.

	wendyMsgString := dht.GpgEncrypt(msgPayload, msgHdr)

	id, err := wendy.NodeIDFromBytes([]byte("FakeID"))
	if err != nil {
		log.Printf("[DHT-BCAST] Error creating fake nodeid %s", err.Error())
	}

	msg := cluster.NewMessage(byte(30), id, []byte(wendyMsgString))

	WendyBroadcast(msg)

	log.Printf("[DHT-BCAST] spreading around XOVER for %s", messageID)

}

func BCastGroup(groupname string) {

	var msgHdr map[string]string
	msgHdr = make(map[string]string)

	msgHdr[TRIBES_H_CMD] = TRIBES_NEWGROUP

	msgPayload := groupname // we are printing it already when requested.

	wendyMsgString := dht.GpgEncrypt(msgPayload, msgHdr)

	id, err := wendy.NodeIDFromBytes([]byte("010203"))
	if err != nil {
		log.Printf("[DHT-BCAST] Error creating fake nodeid %s", err.Error())
	}

	msg := cluster.NewMessage(byte(30), id, []byte(wendyMsgString))

	WendyBroadcast(msg)

	log.Printf("[DHT-BCAST] spreading around GROUP for %s", groupname)

}
