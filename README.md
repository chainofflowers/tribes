# averno
Small, local NNTP server using HTTP  for peering.


I will  write 4 components:

first, the http component:

1. http, to transfer groups, peers and messages.

then the second component, NNTP

1. NNTP server: interface to the user's client and to the peers, in order to exchange messages and groups.

the third module, UpNP
  
>It will try to open port -TO BE DECIDED- on the router for dht.

Due of port 11119, no **root** privileges SHALL be used.

Basically i plan to have 5 threads :

1. Upnp opening the router and keeping it open
2. http looking 
3. nntp listening 
4. one infinite loop checking lifespans and deleting old articles from filesystem.
5. one infinite loop spreading/downloading new messages and groups to/from all known peers, using pastry.

all files will be stored in local ~/news folder. 

Naming convention will be the following: 

<pre>
~/news/posts/x.groupname-YYMMDDHHMMSS-SH1 
</pre>

where:

1. _x_ is "h" for headers, "b" for bodies.
2. _groupname_ is the name of the group (z.b. de.mahl.zeit) , to keep easy  "find all files for a group"
3. _YYMMDDHHMMSS_ is a timestamp, to keep files in natural order and make the NNTP interface easier.
4. _SH1_ is the SH1 hash of the body, to avoid duplicates.
5. One file, one message.

Peers will be saved in 

storing things in such a way will make easier to delete old files by group, 
even manually/cronjob if needed. I am still a  sysadmin :)

I will make it assuming Linux/Unix. If it doesn't works in windows, please restart your PC.

I will choose a set of libraries to use, until they have documentation (**code is not documentation.** If you think code is a good documentation, just work alone) and good examples. Everything which has an incomplete/shitty documentation will be outscoped from the project. 

1. HTTP: [http://tools.ietf.org/html/rfc2616](http://tools.ietf.org/html/rfc2616)[http://tools.ietf.org/html/rfc1945](http://tools.ietf.org/html/rfc1945)
2. NNTP: [https://tools.ietf.org/html/rfc977](https://tools.ietf.org/html/rfc977)
3. UPnP: [https://github.com/prestonTao/upnp](https://github.com/prestonTao/upnp)
4. CONF: [https://github.com/spf13/viper](https://github.com/spf13/viper)
