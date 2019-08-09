# Istio mTLS Spike

## Question

Can we replace the nginx that does mTLS termination in DCS with stuff provided by Istio?

## Answer

Probably. Assuming the app can be changed to parse the `X-Forwarded-Client-Cert` header and the nginx isn't doing anything other that just terminating mTLS.

## Tell me more

We've deployed this to the sandbox cluster in the `dcs-spike-mtls` namespace. I've tried it in GSP local but ended up with the classic `upstream connect error or disconnect/reset before headers. reset reason: connection failure` error. We didn't spend much time looking at it and just did it in a real cluster instead.

Apply the spike YAML:

```
./scripts/gsp-local.sh template && kubectl apply -R -f manifests/gsp-cluster/05-spike-dcs-mtls
```

`curl` the app:

```
curl https://lol.london.sandbox.govsvc.uk -v --cacert mtls-spike/ca-chain.cert.pem --cert mtls-spike/lol.london.sandbox.govsvc.uk.cert.pem --key mtls-spike/lol.london.sandbox.govsvc.uk.key.pem -H"lol: yes m8" --resolve lol.london.sandbox.govsvc.uk:443:$(dig +short nlb.london.sandbox.govsvc.uk | head -n1)
```

You should expect to see the `x-forwarded-client-cert` being set in the response header:

```
* Added lol.london.sandbox.govsvc.uk:443:3.8.231.190 to DNS cache
* Rebuilt URL to: https://lol.london.sandbox.govsvc.uk/
* Hostname lol.london.sandbox.govsvc.uk was found in DNS cache
*   Trying 3.8.231.190...
* TCP_NODELAY set
* Connected to lol.london.sandbox.govsvc.uk (3.8.231.190) port 443 (#0)
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
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256
* ALPN, server accepted to use h2
* Server certificate:
*  subject: C=US; ST=Denial; L=Springfield; O=Dis; CN=lol.london.sandbox.govsvc.uk
*  start date: Jul 26 09:24:00 2019 GMT
*  expire date: Aug  4 09:24:00 2020 GMT
*  common name: lol.london.sandbox.govsvc.uk (matched)
*  issuer: C=US; ST=Denial; O=Dis; CN=lol.london.sandbox.govsvc.uk
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x7ff162009200)
> GET / HTTP/2
> Host: lol.london.sandbox.govsvc.uk
> User-Agent: curl/7.54.0
> Accept: */*
> lol: yes m8
>
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
< HTTP/2 200
< server: istio-envoy
< date: Fri, 09 Aug 2019 14:55:33 GMT
< content-type: application/octet-stream
< content-length: 0
< lol: yes m8
< x-forwarded-client-cert: Hash=ac2be8c512d8e0f41b2a2a89b16fc5f56b935aa8f7577c1e066ce6ca87a2ef3b;Subject="CN=lol.london.sandbox.govsvc.uk,O=Dis,L=Springfield,ST=Denial,C=US";URI=,By=spiffe://cluster.local/ns/spike-dcs-mtls/sa/default;Hash=dd5945ddcc09bd844177ec2cf02f331b10831052231ceee0744c1e45e8aa2740;Subject="";URI=spiffe://cluster.local/ns/istio-system/sa/istio-ingressgateway-service-account
< x-envoy-upstream-service-time: 3
<
* Connection #0 to host lol.london.sandbox.govsvc.uk left intact
```

## tl;dr

get rekt nginx
