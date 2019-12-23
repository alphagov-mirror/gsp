# ADR043: Kubernetes resource access rules

## Status

Pending

## Context

Several different levels of access are required within a kubernetes cluster. The
GSP uses role based access control so these levels are granted to users and
groups via roles.

## Decision

We will create two levels of access within each namespace:

* Operator
* Auditor

The Operator is a relatively permissive read-write role within the namespace.
Developers working on branches that are not part of the release process may be
granted this role in certain namespaces. This is also the role the in-cluster
concourse team for each namespace will be granted.

The Auditor is mostly read-only and will be given to all authenticated users in
the cluster.

The complete list of resource permissions is given in Appendix A.

## Consequences

* Site Reliability Engineers (SRE) will be able to effectively operate the GSP
  while reducing the need to escalate to "admin" level access.
* service team members (e.g. devs) will be able to iterate changes quickly via
  branch-based development and/or `kubectl apply` level access to specific
  namespaces.
* service team members will not be able to interfere with other devs'
  namespaces, "system" namespaces or "prod" namespaces.
* SREs and service team members will not be able to cause service degradations
  or outages to production services.


## Appendix A: Permissions rules

```
ApiGroup: ""
configmaps:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
endpoints:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
events:
  Operator:  get, list, watch
  Auditor:   get, list, watch
limitranges:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
s:
  Operator:  get, list, watch
  Auditor:   get, list, watch
nodes:
  Operator:  get, list, watch
  Auditor:   get, list, watch
persistentvolumeclaims:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   get, list, watch
pods:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
podtemplates:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
resourcequotas:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
secrets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, list, watch
serviceaccounts:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
services:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: "access.govsvc.uk"
principals:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: admissionregistration.k8s.io
mutatingwebhookconfigurations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
validatingwebhookconfigurations:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: apiextensions.k8s.io
customresourcedefinitions:
  Operator:  get, list, watch
  Auditor:   get, list, watch
apiservices:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: apps
daemonsets:
  Operator:  get, list, watch
  Auditor:   get, list, watch
deployments:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
replicasets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
statefulsets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: authentication.istio.io
meshpolicies:
  Operator:  get, list, watch
  Auditor:   get, list, watch
policies:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   get, list, watch

ApiGroup: autoscaling
horizontalpodautoscalers:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: batch
cronjobs:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
jobs:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: bitnami.com
sealedsecrets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: cert-manager.io
certificaterequests:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
certificates:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
clusterissuers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
issuers:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: certificates.k8s.io
certificatesigningrequests:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

pipelines:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
teams:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: config.gatekeeper.sh
configs:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: config.istio.io
adapters:
  Operator:  get, list, watch
  Auditor:   get, list, watch
apikeys:
  Operator:  get, list, watch
  Auditor:   get, list, watch
attributemanifests:
  Operator:  get, list, watch
  Auditor:   get, list, watch
authorizations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
bypasses:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
checknothings:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
circonuses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
cloudwatches:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
deniers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
dogstatsds:
  Operator:  get, list, watch
  Auditor:   get, list, watch
edges:
  Operator:  get, list, watch
  Auditor:   get, list, watch
fluentds:
  Operator:  get, list, watch
  Auditor:   get, list, watch
handlers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
httpapispecbindings:
  Operator:  get, list, watch
  Auditor:   get, list, watch
httpapispecs:
  Operator:  get, list, watch
  Auditor:   get, list, watch
instances:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
kubernetesenvs:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
kuberneteses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
listcheckers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
listentries:
  Operator:  get, list, watch
  Auditor:   get, list, watch
logentries:
  Operator:  get, list, watch
  Auditor:   get, list, watch
memquotas:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
metrics:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
noops:
  Operator:  get, list, watch
  Auditor:   get, list, watch
opas:
  Operator:  get, list, watch
  Auditor:   get, list, watch
prometheuses:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
quotas:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
quotaspecbindings:
  Operator:  get, list, watch
  Auditor:   get, list, watch
quotaspecs:
  Operator:  get, list, watch
  Auditor:   get, list, watch
rbacs:
  Operator:  get, list, watch
  Auditor:   get, list, watch
redisquotas:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
reportnothings:
  Operator:  get, list, watch
  Auditor:   get, list, watch
rules:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
signalfxs:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
solarwindses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
stackdrivers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
statsds:
  Operator:  get, list, watch
  Auditor:   get, list, watch
stdios:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
templates:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
tracespans:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
zipkins:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: coordination.k8s.io
leases:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: crd.k8s.amazonaws.com
eniconfigs:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: crd.projectcalico.org
bgpconfigurations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
bgppeers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
blockaffinities:
  Operator:  get, list, watch
  Auditor:   get, list, watch
clusterinformations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
felixconfigurations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
globalnetworkpolicies:
  Operator:  get, list, watch
  Auditor:   get, list, watch
globalnetworksets:
  Operator:  get, list, watch
  Auditor:   get, list, watch
hostendpoints:
  Operator:  get, list, watch
  Auditor:   get, list, watch
ipamblocks:
  Operator:  get, list, watch
  Auditor:   get, list, watch
ippools:
  Operator:  get, list, watch
  Auditor:   get, list, watch
networkpolicies:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
networksets:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: database.govsvc.uk
postgres:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: events.k8s.io
events:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: extensions
daemonsets:
  Operator:  get, list, watch
  Auditor:   get, list, watch
deployments:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
ingresses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
networkpolicies:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
podsecuritypolicies:
  Operator:  get, list, watch
  Auditor:   get, list, watch
replicasets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: metrics.k8s.io
nodes:
  Operator:  get, list, watch
  Auditor:   get, list, watch
pods:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: monitoring.coreos.com
alertmanagers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
podmonitors:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
prometheuses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
prometheusrules:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
servicemonitors:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: networking.istio.io
destinationrules:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
envoyfilters:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
gateways:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
serviceentries:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
sidecars:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
virtualservices:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: networking.k8s.io
ingresses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
networkpolicies:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: node.k8s.io
runtimeclasses:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: policy
poddisruptionbudgets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
podsecuritypolicies:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: queue.govsvc.uk
sqs:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: rbac.authorization.k8s.io
clusterrolebindings:
  Operator:  get, list, watch
  Auditor:   get, list, watch
clusterroles:
  Operator:  get, list, watch
  Auditor:   get, list, watch
rolebindings:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
roles:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: rbac.istio.io
authorizationpolicies:
  Operator:  get, list, watch
  Auditor:   get, list, watch
clusterrbacconfigs:
  Operator:  get, list, watch
  Auditor:   get, list, watch
rbacconfigs:
  Operator:  get, list, watch
  Auditor:   get, list, watch
servicerolebindings:
  Operator:  get, list, watch
  Auditor:   get, list, watch
serviceroles:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: scheduling.k8s.io
priorityclasses:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: storage.govsvc.uk
s3buckets:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: storage.k8s.io
csidrivers:
  Operator:  get, list, watch
  Auditor:   get, list, watch
csinodes:
  Operator:  get, list, watch
  Auditor:   get, list, watch
storageclasses:
  Operator:  get, list, watch
  Auditor:   get, list, watch
volumeattachments:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: templates.gatekeeper.sh
constrainttemplates:
  Operator:  get, list, watch
  Auditor:   get, list, watch

ApiGroup: verify.gov.uk
certificaterequests:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch
metadata:
  Operator:  create, delete, get, list, patch, update, watch
  Auditor:   delete, get, list, watch

ApiGroup: webhook.cert-manager.io
mutations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
validations:
  Operator:  get, list, watch
  Auditor:   get, list, watch
```
