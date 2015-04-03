package legion

import (
	"os"
	"../tools/"
	"github.com/secondbit/wendy"
  "log"
)

func Initialize() {

  hostname, err := os.Hostname()
  nodename := tools.RandSeq(30)
  networkaddress := tools.ReadIpFromHost()
if err != nil {
  log.Printf("[INFO] %s", "Cannot get own hostname!")
	panic(err.Error())
}
id, err := wendy.NodeIDFromBytes([]byte(nodename+"averno" ))
if err != nil {
  log.Printf("[INFO] %s", "Cannot create a new Node")
	panic(err.Error())
}
node := wendy.NewNode(id,  networkaddress, networkaddress , "AVERNO", 20000)
log.Printf("[INFO] %s", "Nodename: "+nodename)
// this is just to avoid interference with other clusters/dht nodes
credentials := wendy.Passphrase("Chi puote puote, chi non puote se lo scuote")
cluster := wendy.NewCluster(node, credentials)
log.Printf("[INFO] %s", "Cluster initialized")
go func() {
	defer cluster.Stop()
	err := cluster.Listen()
	if err != nil {
		panic(err.Error())
	}
}()
cluster.Join(hostname, 20000) // ports can be different for each Node
log.Printf("[INFO] %s", "Joined Myself")
select {}

}
