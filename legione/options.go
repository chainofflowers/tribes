package legion

import (
	"github.com/secondbit/wendy"
)

// MessagingConfig holds configuration for the messaging component
type MessagingConfig struct {
	// LocalIP is the IP address the node is reachable on within its region
	LocalIP string `short:"i" long:"localip" default:"127.0.0.1" description:"LocalIP is the IP address the node is reachable on within its region"`
	// ExternalIP is the external IP of the node. The node is globally reachable on this IP
	ExternalIP string `short:"e" long:"externalip" default:"127.0.0.1" description:"ExternalIP is the external IP of the node. The node is globally reachable on this IP"`
	// Region is the node's region. Nodes with the same region will be heavily favored by the routing algorithm. Can be omitted.
	Region string `short:"r" long:"region" description:"Region is the node's region. Nodes with the same region will be heavily favored by the routing algorithm. Can be omitted."`
	// Port is the port the messaging componen should listen on
	Port int `short:"p" long:"port" default:"20000" description:"Port is the port the messaging componen should listen on"`
	// ID is the node ID. Must have at least 16 characters (anything over 16 is trimmed). If omitted, a random ID is used
	IDString string `short:"I" long:"id" default:"" description:"ID is the node ID. Must have at least 16 characters (anything over 16 is trimmed). If omitted, a random ID is used"`
	// BootstrapNode is the address and port of a node already in the cluster
	BootstrapNode string `short:"b" long:"bootstrap" description:"BootstrapNode is the address and port of a node already in the cluster. Can be omitted"`
	nodeID        wendy.NodeID
	Verbose       []bool `short:"v" long:"verbose" description:"Verbosity. Can be given multiple times"`
	v             int    // verbosity level as a number
}
