package legion

import (
	"../tools/"
	"github.com/secondbit/wendy"
	"log"
	"os"
)

func Initialize() {

	var options MessagingConfig

	log.Printf("[INFO] %s", "Baking a nice Pastry cluster")

	options.nodeID = randomID()
	options.IDString = options.nodeID.String()
	options.ExternalIP = tools.ReadIpFromHost()
	options.LocalIP = "127.0.0.1"
	options.Region = "AVERNO"
	options.Port = GetClusterPort()

	node := wendy.NewNode(options.nodeID, options.LocalIP, options.ExternalIP, options.Region, options.Port)
	log.Printf("[INFO] NodeName : %s", options.nodeID)

	// this is just to avoid interference+bullshit with other clusters/dht nodes
	credentials := wendy.Passphrase("Chi puote puote, chi non puote se lo scuote")
	// actually no secret credentials are needed

	cluster := wendy.NewCluster(node, credentials)
	cluster.SetHeartbeatFrequency(10)
	cluster.SetNetworkTimeout(1)
	cluster.SetLogLevel(wendy.LogLevelWarn)

	cluster.RegisterCallback(&debugWendy{})

	log.Printf("[INFO] %s", "Cluster initialized, now start listening")

	go func() {
		defer cluster.Stop()
		err := cluster.Listen()
		if err != nil {
			log.Printf("[WTF] Cannot listen on port %s SYSADMIN!", options.Port)
			os.Exit(1)
		}
	}()

	// err = cluster.Join("127.0.0.1", 20000)
	// if err != nil {
	// 	log.Printf("[OOPS] Cannot join the cluster %s", err)
	// }

	// log.Printf("[INFO] %s", "Joined Myself, now looking for other nodes")
	select {}

}
