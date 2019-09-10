# Public DNS

The GSP allows applications to set DNS records (route53) based on istio `Gateway` resources on a per-namespace basis. The `Gateway` resource needs to carry the correct annotation to ensure the `external-dns` instance picks it up.

To set the DNS entry for the [gsp-canary][]:

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  annotations:
    externaldns.k8s.io/namespace: sandbox-main
  name: sandbox-gsp-canary-ingress
  namespace: sandbox-main
spec:
  selector:
    istio: sandbox-main-ingressgateway
  servers:
  - hosts:
    - canary.london.sandbox.govsvc.uk
    port:
      name: https
      number: 443
      protocol: HTTPS
    tls:
      mode: SIMPLE
      privateKey: /etc/istio/ingressgateway-certs/tls.key
      serverCertificate: /etc/istio/ingressgateway-certs/tls.crt
```

Each namespace has an instance of [external-dns][] running that will configure DNS A records to point at the load balancer created for the ingressgateway in the same namespace. The `externaldns.k8s.io/namespace: sandbox-main` annotation ensures the `external-dns` instance in the namespace picks it up.

The locations of the TLS certificates will depend on the installation context. This example above uses the per-namespace gateway locations.

[external-dns]: https://github.com/kubernetes-incubator/external-dns
[gsp-canary]: https://github.com/alphagov/gsp/tree/master/components/canary
