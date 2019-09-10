# Public TLS Certificates

TLS certificates can optionally be provisioned via [cert-manager][] in GSP.

By default, the GSP has a `ClusterIssuer` named `letsencrypt-r53` that is configured to provision TLS certificates supplied by [LetsEncrypt][] via the DNS01 ACME challenge. For example: to provision a TLS certificate for the [gsp-canary][] in the sandbox cluster that is mounted into its local istio-ingressgateway:

```yaml
apiVersion: certmanager.k8s.io/v1alpha1
kind: Certificate
metadata:
  name: ingress-certificate
  namespace: sandbox-main
spec:
  acme:
    config:
    - dns01:
        provider: route53
      domains:
      - canary.london.sandbox.govsvc.uk
  dnsNames:
  - canary.london.sandbox.govsvc.uk
  issuerRef:
    kind: ClusterIssuer
    name: letsencrypt-r53
  secretName: istio-ingressgateway-certs
```

To adapt the above example for use elsewhere, just change `.spec.acme.config[].domains` and `.spec.dnsNames`.

> **Note:** cert-manager will need to be able to modify the DNS of the domains listed in the certificate in order to perform the DNS challenge. At the time of writing that only applies to the cluster domain.

If you're following the per-namespace gateway pattern the certificate needs to include all the domains and host names for the namespace as the gateway can only mount a single secret (and so should be configured via the `-cluster-config` chain rather than an application chart).

[cert-manager]: https://docs.cert-manager.io/en/latest/
[gsp-canary]: https://github.com/alphagov/gsp/tree/master/components/canary
[LetsEncrypt]: https://letsencrypt.org/
