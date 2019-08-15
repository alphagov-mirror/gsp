# nginx mTLS Spike

## Question

Can we use nginx to terminate mTLS whilst still having a single istio ingress?

## Answer

Probably.

## Tell me more

We've deployed this to the sandbox cluster in the `dcs-spike-mtls` namespace.

Apply the spike YAML:

```
./scripts/gsp-local.sh template && kubectl apply -R -f manifests/gsp-cluster/05-spike-dcs-mtls
```

`curl` the app:

```
curl https://lol.london.sandbox.govsvc.uk -v --cacert mtls-spike/ca-chain.cert.pem --cert mtls-spike/lol.london.sandbox.govsvc.uk.cert.pem --key mtls-spike/lol.london.sandbox.govsvc.uk.key.pem -H"lol: yes m8" --resolve lol.london.sandbox.govsvc.uk:443:$(dig +short nlb.london.sandbox.govsvc.uk | head -n1)
* Added lol.london.sandbox.govsvc.uk:443:3.9.121.12 to DNS cache
* Rebuilt URL to: https://lol.london.sandbox.govsvc.uk/
* Hostname lol.london.sandbox.govsvc.uk was found in DNS cache
*   Trying 3.9.121.12...
* TCP_NODELAY set
* Connected to lol.london.sandbox.govsvc.uk (3.9.121.12) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* Cipher selection: ALL:!EXPORT:!EXPORT40:!EXPORT56:!aNULL:!LOW:!RC4:@STRENGTH
* successfully set certificate verify locations:
*   CAfile: mtls-spike/ca-chain.cert.pem
  CApath: none
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Request CERT (13):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Certificate (11):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS handshake, CERT verify (15):
* TLSv1.2 (OUT), TLS change cipher, Client hello (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS change cipher, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES256-GCM-SHA384
* ALPN, server accepted to use http/1.1
* Server certificate:
*  subject: C=US; ST=Denial; L=Springfield; O=Dis; CN=lol.london.sandbox.govsvc.uk
*  start date: Jul 26 09:24:00 2019 GMT
*  expire date: Aug  4 09:24:00 2020 GMT
*  common name: lol.london.sandbox.govsvc.uk (matched)
*  issuer: C=US; ST=Denial; O=Dis; CN=lol.london.sandbox.govsvc.uk
*  SSL certificate verify ok.
> GET / HTTP/1.1
> Host: lol.london.sandbox.govsvc.uk
> User-Agent: curl/7.54.0
> Accept: */*
> lol: yes m8
>
< HTTP/1.1 200 OK
< Server: nginx/1.17.2
< Date: Thu, 15 Aug 2019 09:57:12 GMT
< Content-Type: application/octet-stream
< Content-Length: 0
< Connection: keep-alive
< lol: yes m8
<
* Connection #0 to host lol.london.sandbox.govsvc.uk left intact
```
