# Tribes
Small, local NNTP server using P2P  for peering.


I will  write several components:

first, the http component:

1. http/whatever, to transfer groups, peers and messages.

then the second component, NNTP

1. NNTP server: interface to the user's client and to the peers, in order to exchange messages and groups.

the third module, UpNP
  
It will try to open port -Configurable- on the router for http.

Thanks to port 11119, no **root** privileges SHALL be used.

Basically i plan to have 5 threads :

1. Upnp opening the router and keeping it open
2. http/whatever for distributing
3. nntp listening 
4. one infinite loop checking lifespans and deleting old articles from filesystem.
5. one infinite loop spreading/downloading new messages and groups to/from all known peers, using P2P.

The working folder will be stored in local ~/News folder. 

Naming convention for messages will be the following: 

<pre>
~/News/messages/y-groupname-SERIAL-UUID@TIMESTAMP
</pre>


where:

1. _y_ is "h" for headers, "b" for bodies, "x" for XOVER database
2. _groupname_ is the name of the group (z.b. de.mahl.zeit) , no "-" char allowed.
3. _SERIAL_ is a unique identifier of the message, by group: 6 decimal digits.
4. _UUID_ is a unique identifier of the message in general, to avoid duplicates: 38 chars [a-z][A-Z].
6. _@_ to respect RFC850 
7. _TIMESTAMP_ is time of submission in format: YYMMDDHHMM
8. One file, one message.

Please notice: being the '-' a separator, it won't be possible to have groups with '-' in the name.


Peers will be saved in 

<pre>
~/news/peers/peers.initial  // the first host to connect to download other peers.
~/news/peers/peers.active   // the peers we can reach. This is the list to be shared
~/news/peers/peers.all      // all the peers we know from others
</pre>

A worker thread will keep updated the "active" list. Also it will prune the "all" list.

Newsgroup will be stored as:

<pre>
~/news/groups/ng.local  // groups which are created locally. Always considered "new". To be exposed to peers
~/news/groups/ng.active   // groups which are subscribed by the local client.
~/news/groups/ng.all      //all the groups we know. 
</pre>

a running task will take care of pruning groups.all. Pruning means "making it equal to groups.active". 

Storing informations in such a way will make easier to delete old files by group 
Also  manual/cronjob deletion if needed. I am still a  sysadmin :)

I am  assuming it runs on Linux/Unix. If it doesn't works in windows, please restart your PC.

I will choose a set of libraries to use, if/when needed:

1. Until they have documentation (**code is not documentation.** If you think code is a good documentation, just work alone) 
2. Until they have good examples. 
3. Until the code makes sense.

Everything which has an incomplete/shitty documentation will be outscoped.

1. HTTP: [http://tools.ietf.org/html/rfc2616](http://tools.ietf.org/html/rfc2616)[http://tools.ietf.org/html/rfc1945](http://tools.ietf.org/html/rfc1945)
2. NNTP: [https://tools.ietf.org/html/rfc977](https://tools.ietf.org/html/rfc977) - IN PROGRESS -
3. UPnP: [https://github.com/prestonTao/upnp](https://github.com/prestonTao/upnp) - **DONE** -
4. CONF: [https://github.com/spf13/viper](https://github.com/spf13/viper) - **DONE** - 


To compile it:

1. Download the code.
2. Enter the folder tribes
3. go get ; go build

done.
