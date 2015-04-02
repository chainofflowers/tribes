# averno
Small, local NNTP server using HDT as a peering mechanism.


I will  write 4 components:

## first, the dht component:##

>dht : using the dht library, the server will be able to find other peers.

##then the second component, NNTP## 

>-NNTP client: once another peer is found, using NNTP it will upload new contents and download from the peer.
-NNTP server: interface to the user's client and to the peers, in order to exchange messages and groups.

#the third module, UpNP
  
>It will try to open port 11119 on the router (nntp) and some UDP port for dht.


Basically i plan to have 4 threads :

>-Upnp opening the router and keeping it open
-hdt looking for peers
-nntp listening
-an infinite loop spreading/downloading new messages and groups to all the known peers, using NNTP client.

