# Per-namespace istio gateways

We are migrating from using a single `ingressgateway` in `istio-system` to an `ingressgateway` in each namespace. You must opt into / enable an ingressgateway in each namespace via `*-cluster-config` values.

For example, from `tech-ops-cluster-config`'s `sandbox-values`:

```yaml
namespaces:
- name: sandbox-connector-node-metadata
  owner: alphagov
  repository: verify-metadata
  branch: sandbox
  path: ci/sandbox
  requiredApprovalCount: 0
```

To opt-in to a gateway for ingress:

```yaml
namespaces:
- name: sandbox-connector-node-metadata
  owner: alphagov
  repository: verify-metadata
  branch: sandbox
  path: ci/sandbox
  requiredApprovalCount: 0
  ingress:
    enabled: true
```

To opt-in and expose additional ports (`3306` for MySQL in this example):

```yaml
namespaces:
- name: sandbox-connector-node-metadata
  owner: alphagov
  repository: verify-metadata
  branch: sandbox
  path: ci/sandbox
  requiredApprovalCount: 0
  ingress:
    enabled: true
    ports:
    - port: 3306
      name: tcp-mysql
      targetPort: 3306
```

This will create a `Service` in the namespace corresponding to an istio envoy gateway of type `LoadBalancer`. This in turn causes the EKS control plane to provision a network load balancer (NLB) for the service.
