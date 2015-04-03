package legion

import (
	"os"
	"../tools/"
	"github.com/secondbit/wendy"
  "log"
)

func Initialize() {

  log.Printf("[INFO] %s", "Baking a nice Pastry cluster")

  hostname, err := os.Hostname()
  nodename := tools.RandSeq(30)
  networkaddress := tools.ReadIpFromHost()
if err != nil {
  log.Printf("[WTF] %s", "Cannot get own hostname! SYSADMIN!")
	panic(err.Error())
}
id, err := wendy.NodeIDFromBytes([]byte(nodename+"averno" ))
if err != nil {
  log.Printf("[INFO] %s", "Cannot create a new Node because of:")
	panic(err.Error())
}
node := wendy.NewNode(id,  networkaddress, networkaddress , "AVERNO", 20000)
log.Printf("[INFO] %s", "Nodename: "+nodename)

// this is just to avoid interference+bullshit with other clusters/dht nodes
credentials := wendy.Passphrase("Chi puote puote, chi non puote se lo scuote")
// actually no secret credentials are needed

cluster := wendy.NewCluster(node, credentials)
log.Printf("[INFO] %s", "Cluster initialized, now start listening")

go func() {
	defer cluster.Stop()
	err := cluster.Listen()
	if err != nil {
    log.Printf("[WTF] %s", "Cannot listen on port 20000. SYSADMIN!")
		panic(err.Error())
	}
}()


cluster.Join(hostname, 20000) // we choose 20000 
log.Printf("[INFO] %s", "Joined Myself, now looking for other nodes")
select {}


}
