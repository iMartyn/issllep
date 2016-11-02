iSSLLEp
=======
**i** nternal **SSL** **L** ets **E** ncrypt **p** roxy
-------------------------------------

This little proxy should enable internal-only SSL certs to be created.  The idea being that issllep sits on a box responding to LetsEncrypt(LE)  Queries, but your actual Web Server is inaccessible to LE.

This is great for internal-only environments such as test or staging environments where LE should not be able to get to them.

The intention is that a wildcard DNS be set up externally but that you can somehow resolve names internally fully.  Example: `www.mystagingdomain.com` internally could resolve to `172.16.1.1` but externally the only record would be `*.mystagingdomain.com` which would resolve to the external IP of your box running this proxy server.

The proxy on first run (actually on first migration) will generate a pre-shared key for you to generate SSL certs with.  The intent in version 1.0 is to use SSL private key auth, but that will take some trickery to get it working and automatable.
