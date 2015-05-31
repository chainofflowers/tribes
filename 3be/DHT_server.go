package tribe

import (
	"github.com/secondbit/wendy"
	"log"
	"tribes/config"
	"tribes/tools"
)

type WendyApplication struct {
}

var (
	cluster *wendy.Cluster
	id      wendy.NodeID
	node    *wendy.Node
	err     error
	cred    wendy.Credentials
)

func init() {

	TribeID := config.GetTribeID()
	WendyID := tools.RandSeq(42)
	log.Printf("[DHT] Volatile node ID: %s", WendyID)

	id, err = wendy.NodeIDFromBytes([]byte(WendyID))
	if err != nil {
		log.Printf("[DHT] Can't create the NodeID: %s", WendyID)
	}

	node = wendy.NewNode(id, tools.ReadIpFromHost(), tools.ReadIpFromHost(), "Tribes", config.GetClusterPort())
	log.Printf("[DHT] Node created")

	cred = wendy.Passphrase(TribeID)

	cluster = wendy.NewCluster(node, cred)
	log.Printf("[DHT] Cluster initialized")

	go cluster.Listen()
	log.Printf("[DHT] Listening")

	if tmp_boot := config.GetBootStrapHost(); tmp_boot != "127.0.0.1" {
		tmp_port := config.GetBootStrapPort()
		cluster.Join(tmp_boot, tmp_port)
		log.Printf("[DHT] Trying to join cluster at %s:%d", tmp_boot, tmp_port)
	}

	app := &WendyApplication{}
	cluster.RegisterCallback(app)
	log.Printf("[DHT] Engine functional ")

}

func Tribe_Engine_Start() {
	log.Printf("[DHT] DHT Engine Exists")

}

func (app *WendyApplication) OnError(err error) {
	log.Printf("[DHT] OOPS: %s", err.Error())

}

func (app *WendyApplication) OnDeliver(msg wendy.Message) {
	log.Printf("[DHT] Received message: %s", msg.String())
	var mypayload TribePayload

	mypayload.TPbuffer = []byte(msg.String())
	mypayload.TPsize = len(mypayload.TPbuffer)
	Tribes_Interpreter(mypayload)
}

func (app *WendyApplication) OnForward(msg *wendy.Message, next wendy.NodeID) bool {
	log.Printf("[DHT] Forwarding message %s to node %s.", msg.Key, next)
	return true // return false if you don't want the message forwarded
}

func (app *WendyApplication) OnNewLeaves(leaves []*wendy.Node) {
	log.Println("[DHT] New leaves: ", leaves)
}

func (app *WendyApplication) OnNodeJoin(node wendy.Node) {
	log.Println("[DHT] Node joined: ", node.ID)
}

func (app *WendyApplication) OnNodeExit(node wendy.Node) {
	log.Println("[DHT] Node left: ", node.ID)
}

func (app *WendyApplication) OnHeartbeat(node wendy.Node) {
	log.Println("[DHT] Received heartbeat from ", node.ID)
}
