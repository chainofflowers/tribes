package tribe

import (
	"github.com/secondbit/wendy"
	"log"
	"math"
	"tribes/config"
	"tribes/tools"
)

type WendyApplication struct {
}

var (
	cluster  *wendy.Cluster
	id       wendy.NodeID
	mynode   *wendy.Node
	err      error
	cred     wendy.Credentials
	AllNodes map[wendy.NodeID]string
)

func init() {

	AllNodes = make(map[wendy.NodeID]string)

	TribeID := config.GetTribeID()
	WendyID := tools.RandSeq(16)
	log.Printf("[DHT] Volatile node ID: %s", WendyID)

	id, err = wendy.NodeIDFromBytes([]byte(WendyID))
	if err != nil {
		log.Printf("[DHT] Can't create the NodeID: %s", WendyID)
	}

	mynode = wendy.NewNode(id, tools.ReadIpFromHost(), tools.ReadIpFromHost(), "Tribes", config.GetClusterPort())
	log.Printf("[DHT] Node created")

	cred = wendy.Passphrase(TribeID)

	cluster = wendy.NewCluster(mynode, cred)
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

	cluster.SetHeartbeatFrequency(5)
	cluster.SetNetworkTimeout(300)

}

func Tribe_Engine_Start() {
	log.Printf("[DHT] DHT Engine Exists")

}

// we don't like crashes, so we just log it.
func (app *WendyApplication) OnError(err error) {
	log.Printf("[DHT] OOPS: %s", err.Error())

}

func (app *WendyApplication) OnDeliver(msg wendy.Message) {
	log.Printf("[DHT] Received message: %s", msg.String())
	var mypayload TribePayload

	mypayload.TPbuffer = []byte(msg.String())
	mypayload.TPsize = len(mypayload.TPbuffer)
	Tribes_Interpreter(mypayload)

	// we forward with a lesser TTL
	if msg.Purpose > 30 {
		newpurpose := msg.Purpose - 1
		AnyCastSpread(newpurpose, msg, cluster)
	}

}

// just let it forward
func (app *WendyApplication) OnForward(msg *wendy.Message, next wendy.NodeID) bool {
	log.Printf("[DHT] Forwarding message %s to node %s.", msg.Key, next)
	return true // return false if you don't want the message forwarded
}

func (app *WendyApplication) OnNewLeaves(leaves []*wendy.Node) {

	log.Println("[DHT] New leaves: ", leaves)

	for wItem := range leaves {

		nID := leaves[wItem]

		AllNodes[nID.ID] = "active"
		log.Printf("[DHT]Leave %s added", nID.ID)
	}

}

// add the node we know entered
func (app *WendyApplication) OnNodeJoin(node wendy.Node) {
	AllNodes[node.ID] = "active"
	log.Println("[DHT] Node joined: ", node.ID)
}

// remove the node which exited
func (app *WendyApplication) OnNodeExit(node wendy.Node) {
	delete(AllNodes, node.ID)
	log.Println("[DHT] Node left: ", node.ID)
}

// if we receive an heartbit, we know this node is active
func (app *WendyApplication) OnHeartbeat(node wendy.Node) {
	AllNodes[node.ID] = "active"
	log.Println("[DHT] Received heartbeat from ", node.ID)
}

//We will use something similar to AnyCast from IPv6 to spread the messages around.
//Since Wendy says we can only use n > 16 as a "purpose", I will use the purpose as TTL
//So the purpose of 30 will be TTL = 0. 31 will be TTL=1 , 32 will be TTL =2 , and so on.
//the idea is "each node will advertise each other known nodes about a new message,
//until the TTL will expire. Given a separation layer 6 (globally), TTL=10 is overkill.
func AnyCastSpread(TTL uint8, mymessage wendy.Message, mycluster *wendy.Cluster) {

	for nID, _ := range AllNodes {
		if nID != mynode.ID {
			msg := mycluster.NewMessage(TTL, nID, []byte(mymessage.String()))
			err := mycluster.Send(msg)
			if err != nil {
				log.Println("[DHT] Can't send a message to: ", nID)
			}
		}
	}

}

// This is how we initiate a broadcast. We choose the TTL as log2 of the amount of machines
// into the cluster. Then we spread it around using AnyCastSpread.
func WendyBroadcast(message wendy.Message) {

	var myTTL uint8
	myTTL = 1
	nodeNum := float64(len(AllNodes))
	if ll := math.Log2(nodeNum); ll >= 1 {
		myTTL = 30 + uint8(ll)
		AnyCastSpread(byte(myTTL), message, cluster)
	} else {
		log.Println("[DHT] Only node in the cluster, nothing to do")
	}

}
