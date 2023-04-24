# HTTP C2 IOCs


## Port Portocol Mix Match

In wireshark find the traffic using the "http" display filter.

Now look at the tcp port associated with the traffic. 443 is most commonly used for HTTPS, so this should be TLS traffic.


## Header & Site Mix Match

Stil in Wireshark, right click and follow the TCP or HTTP stream.

At the top of the content, you can see the headers.  The headers for client requests have an Authority header that says "Microsoft.com" yet it's for sure not going to Microsoft.

But the client header information could be unique per client, it would be better to search your whole org for IOCs related to this server, regardless of IP. 

The server header has the "Mode" header, this could be used for signatures.

What else? Share with the class.

Did any security tools automatically identify the traffic or exe as malicous on stock Windows?  Will yours?

[Become The Threat Home](../../../RSAC23-Become-The-Threat/README.md)