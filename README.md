# averno
Small, local NNTP server using HDT as a peering mechanism.


I will  write 4 components:

first, the dht component:

>-dht : using the dht library, the server will be able to find other peers.

then the second component, NNTP

1. NNTP client: once another peer is found, using NNTP it will upload new contents and download from the peer.
2. NNTP server: interface to the user's client and to the peers, in order to exchange messages and groups.

the third module, UpNP
  
>It will try to open port 11119 on the router (nntp) and some UDP port for dht.


Basically i plan to have 5 threads :

1. Upnp opening the router and keeping it open
2. dht looking for peers
3. nntp listening 
4. one infinite loop checking lifespans and deleting old articles from filesystem.
5. one infinite loop spreading/downloading new messages and groups to/from all known peers, using NNTP client.

all files will be stored in local ~/News folder. Since of port 11119 , no **root** privileges SHALL be needed.





