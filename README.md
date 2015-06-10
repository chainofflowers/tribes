# Tribes
Small, local NNTP server **using P2P/DHT**  for peering.

The aim of Tribes is to have a small, zero-config (or little-config) personal NNTP
server. This is for personal use, like running on your laptop,  your PC at home, and so on.
This is why I call it "_personal server_".

Each _personal server_ will use DHT [Pastry](http://en.wikipedia.org/wiki/Pastry_%28DHT%29)
to exchange messages and newsgroups. Also it supports NAT-UpNP.

Also, Tribes allows each node to create newsgroups. This means editing a text file
is enough.

To avoid spam and crimes, Tribes servers are grouped into "tribes" . A "Tribe" is nothing
different than a passphrase, which needs to be the same for all the member of the tribe.
From now on, we name this passphrase as "TribeID"

The minimal lenght of the passphrase is 32 chars.

With no proper key, a peer cannot join a tribe. Groups and folders are exchanged
after encryption with PGP, using the TribeID as a key.

All files are stored into ${HOME}/News. 

## How it works 

Like any DHT network, it needs an "initial" (AKA "bootstrap") node. Your node
will connect to the initial node, and it will join the DHT network. It will get
the DHT cache, and now you are in.

**NOTES FOR THE BETA RELEASE** you will receive any message posted **after** you joined
the cluster. The "full alignment" will come later.

You need to set your TribeID in the config file, and also the TCP port you want to use.

If your router is not supporting NAT-UpNP, you''ll need to configure it accordingly.

## How to use it.

Download the executable (or build it) and start it once.
On linux just type ""./tribes" , on Windows just open your terminal and do the same.

Then shut it  down, with CONTROL-C.

This will create the folders and a default config file.

You can find the configuration file in ${HOME}/News/config.toml

Just open it with your editor. You will see:

<pre>
TLSPORT = "21000"
MyTribeID = "AdzfNdsMAajMMuPpVsNXvWWxIDohwppz"
MyPublicHost = "whatever.example.com"
MyBootStrapHost = "127.0.0.1"
MyBootStrapPort = "21000"
</pre>

**TLSPORT** is the port you want to use in your PC. This is the port which
will be advertised to the home router (if one) via NAT-UpNP. If your router
is not NAT-UpNP ready, you need to forward this port manually.

**MyTribeID** You need this to enter the tribe. Either you get this from other members
of an existing tribe, or you want to create your own tribe. If you want to create your
own tribe, just invent a string. It has to be longer than 32 chars.

**MyPublicHost** Currently not used, but planned for a future release.You can leave it 
like it is.

**MyBootStrapHost** The IP of the host you want to use as a bootstrap. beta0.1/2 only
supports IP addresses. I will add DNS resolution soon. If you want to be the bootstrap
node, just set as "127.0.0.1"

**MyBootStrapPort** the port your bootstrap server is listening at. 

Once you configured , just save the file and restart tribes using the command line.

Now you can open the Newsreader you like , and set a newsserver as "127.0.0.1:11119"
No autentication is needed.

In case you want to configure your own newsgroups, just open the file ${HOME}/News/groups/ng.local
and write  a list of them like:

<pre>
it.stinks
it.stinks.more
it.stinks.even.more
</pre>

of yourse, you can put there the names you like. This **example** setup will create 3 groups,
named like "it.stinks", "it.stinks.more", "it.stinks.even.more". 

# FAQ

- Q. Do I need to connect to some ISP's news server?
- A. No. The concept is to have a NNTP service with NO big servers involved. Just P2P.

- Q. May I use it as local NNTP server on my huge machine?
- A. Well, it only listens on localhost. Maybe for local users? Dunno. Up to you.

- Q. Should I expose port 11119 to the internet?
- A. No.I put little care about security when writing the _beloved_ NNTP interface.

- Q. But I want to use Tribes from the internet. Is there a way you suggest? haproxy, nginx?
- A. If you like, **don't** just redirect tcp ports. Use a real Layer7 proxy, like leafnode. Still I think you should not.

- Q. C'mon! What can happen if I expose the NNTP interface?
- A. Spammers will flood you and your "tribe"'s member with bullshit: _there is no autentication_.

- Q. Is it a darknet?
- A. No. Encryption has the only goal of keeping people out of a tribe _until they can't get the TribeID_. As far I see, **pastry implementation is in clear text**.

- Q. Does it grants my privacy?
- A. USENET is a **public** space. When you are in public, **people can see you**.

- Q. When you will finish it?
- A. Dunno. I still have lot of ideas about it. So I will change the storage later, with a more efficient one. I will add a configuration web interface. There are many ways to do that. So I am still collecting ideas.

- Q. So what's the plan?
- A. Now we have beta0.2 . I will continue with beta0.X until I am sure "it just works" and the network protocol is consistent. Then I will do Release 1.0. After of this, I will add  everything I like.

- Q. May I help? How?
- A. You could do testing, joining a test cluster. In such a case , you will be requested to join the Tribe , join discussions, and so on. In case you see a bug, just collect the logs and then send them to me. Logs are in ${HOME}\News\logs\tribes.2015-Jun-10.1700.log

- Q. How it works this "_sent them to me_"?
- A. Just write _gitadmin AT wolfstep.cc_ and ask to partecipate. Then we can use the GitBucket to open issues.

- Q. Could I contribute to the code, also?
- A. Being honest, no. This is something I do for fun, to experiment with golang, to implement stuffs. Don't take me wrong if I say no: is not about you. It is _my_ fun. :)

- Q. Do you think it will be a success?
- A. Indeed, it is: I had lot of fun writing it. This was the goal since the beginning: I'm  not a PRO developer,  this is just my hobby.

- Q. Imagine I create a "tribe" in my local net, and then I expose the NNTP port. Is this a replicated cluster of NNTP servers?
- A. <u>Tribes was not designed for production use</u>. Btw, still is a _beta_. Functionally speaking: yes, you can get a replicated cluster of NNTP servers like this.

- Q. I want to join the cluster, I have credentials, but the Bootstrap node is often slow. What can I do?
- A. You can use any node to bootstrap. Until the node is part of a tribe, and you can reach it, it is a "bootstrap node". This is by design of DHT. There is no "designated" first node.

- Q. I want to start a tribe.
- A. Just decide your passphrase, and give to your friend your public IP, your port, and the TribeID itself. 

- Q. May I give my friends the whole "config.toml" ready to use? My friends are not good in configurations.
- A. Absolutely yes. You can do even more: just put the tribes executable into ${HOME}/News, put the configuration you want to offer, and zip the folder. Your friends will just need to unpack the "News" folder into ${HOME} and start tribes from there.

- Q. You implemented just some pieces of the NNTP protocol.
- A. NNTP protocol was designed by a barrel of drunken monkeys (Or something like that). All clients are behaving differently, and there is no server on the planet which is 100% compliant with it. Too much shit, I think. During the testing phase with MacSoup, Thunderbird, Pan , Slrn, leafnode,  I decided to implement a reasonable subset of it. Like everyone is doing.

- Q. The newsreader i love is not working with that.
- A. Please send me the logs and I will see which spaghetti-NNTP-feature your client is trying to use. I will eventually add it,  I cannot say when.


