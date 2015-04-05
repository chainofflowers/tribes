package legion

import (
	"github.com/secondbit/wendy"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomID() (id wendy.NodeID) {
	id[0] = uint64(uint64(rand.Uint32())<<32 | uint64(rand.Uint32()))
	id[1] = uint64(uint64(rand.Uint32())<<32 | uint64(rand.Uint32()))
	return
}

type debugWendy struct {
}

func (app *debugWendy) OnError(err error) {
	log.Printf("[OMG] Incoming Bullshit : %s", err)
}

func (app *debugWendy) OnDeliver(msg wendy.Message) {
	log.Print("Received message: ", msg)
}

func (app *debugWendy) OnForward(msg *wendy.Message, next wendy.NodeID) bool {
	log.Printf("Forwarding message %s to Node %s.", msg.Key, next)
	return true 
    // return false if you don't want the message forwarded
}

func (app *debugWendy) OnNewLeaves(leaves []*wendy.Node) {
	log.Print("Leaf set changed: ", leaves)
}

func (app *debugWendy) OnNodeJoin(node wendy.Node) {
	log.Print("Node joined: ", node.ID)
}

func (app *debugWendy) OnNodeExit(node wendy.Node) {
	log.Print("Node left: ", node.ID)
}

func (app *debugWendy) OnHeartbeat(node wendy.Node) {
	log.Print("Received heartbeat from ", node.ID)
}
