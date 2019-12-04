# ADR043: Kubernetes resource access rules

## Status

Pending

## Context

Several different levels of access are required within a kubernetes cluster. The
GSP uses role based access control so these levels are granted to users and
groups via roles.

## Decision

We will grant users access to the resources according to the rules given below
(see Appendix A). Namespaces will be split into "prod" and "non-prod" as the
access for some resources will be different. Here "prod" and "non-prod" refer to
classifications of namespace. Examples of "prod" namespaces:

- kube-system
- istio-system
- gsp-system
- verify-doc-checking-build
- verify-doc-checking-prod
- verify-proxy-node-build
- verify-proxy-node-prod

The "non-prod" namespaces are those created for developers to rapidly iterate on
changes to test they work in a GSP cluster environment and fall completely
outside the path to release.

Also included in the permissions rules below are those given to the in-cluster
concourse. These permissions apply to both "prod" and "non-prod" namespaces.

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
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
endpoints:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
events:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
limitranges:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
namespaces:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
nodes:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
persistentvolumeclaims:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      get, list, watch
  concourse: create, delete, get, list, patch, update, watch
pods:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
podtemplates:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
resourcequotas:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
secrets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, list, watch
  concourse: create, delete, get, list, patch, update, watch
serviceaccounts:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
services:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: "access.govsvc.uk"
principals:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: admissionregistration.k8s.io
mutatingwebhookconfigurations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
validatingwebhookconfigurations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: apiextensions.k8s.io
customresourcedefinitions:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
apiservices:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: apps
daemonsets:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
deployments:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
replicasets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
statefulsets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: authentication.istio.io
meshpolicies:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
policies:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: autoscaling
horizontalpodautoscalers:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: batch
cronjobs:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
jobs:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: bitnami.com
sealedsecrets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: cert-manager.io
certificaterequests:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
certificates:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
clusterissuers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
issuers:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: certificates.k8s.io
certificatesigningrequests:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: concourse.k8s.io
pipelines:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
teams:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: config.gatekeeper.sh
configs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

----------- DP below this line

ApiGroup: config.istio.io
adapters:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
apikeys:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
attributemanifests:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
authorizations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
bypasses:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
checknothings:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
circonuses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
cloudwatches:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
deniers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
dogstatsds:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
edges:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
fluentds:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
handlers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
httpapispecbindings:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
httpapispecs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
instances:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
kubernetesenvs:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
kuberneteses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
listcheckers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
listentries:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
logentries:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
memquotas:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
metrics:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
noops:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
opas:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
prometheuses:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
quotas:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
quotaspecbindings:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
quotaspecs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
rbacs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
redisquotas:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
reportnothings:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
rules:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
signalfxs:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
solarwindses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
stackdrivers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
statsds:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
stdios:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
templates:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
tracespans:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
zipkins:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: coordination.k8s.io
leases:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: crd.k8s.amazonaws.com
eniconfigs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: crd.projectcalico.org
bgpconfigurations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
bgppeers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
blockaffinities:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
clusterinformations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
felixconfigurations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
globalnetworkpolicies:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
globalnetworksets:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
hostendpoints:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
ipamblocks:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
ippools:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
networkpolicies:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
networksets:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: database.govsvc.uk
postgres:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: events.k8s.io
events:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: extensions
daemonsets:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
deployments:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
ingresses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
networkpolicies:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
podsecuritypolicies:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
replicasets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: metrics.k8s.io
nodes:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
pods:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: monitoring.coreos.com
alertmanagers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
podmonitors:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
prometheuses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
prometheusrules:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
servicemonitors:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: networking.istio.io
destinationrules:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
envoyfilters:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
gateways:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
serviceentries:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
sidecars:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
virtualservices:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: networking.k8s.io
ingresses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
networkpolicies:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: node.k8s.io
runtimeclasses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: policy
poddisruptionbudgets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
podsecuritypolicies:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: queue.govsvc.uk
sqs:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: rbac.authorization.k8s.io
clusterrolebindings:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
clusterroles:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
rolebindings:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
roles:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: rbac.istio.io
authorizationpolicies:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
clusterrbacconfigs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
rbacconfigs:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
servicerolebindings:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
serviceroles:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: scheduling.k8s.io
priorityclasses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: storage.govsvc.uk
s3buckets:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: storage.k8s.io
csidrivers:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
csinodes:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
storageclasses:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
volumeattachments:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: templates.gatekeeper.sh
constrainttemplates:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch

ApiGroup: verify.gov.uk
certificaterequests:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch
metadata:
  non-prod:  create, delete, get, list, patch, update, watch
  prod:      delete, get, list, watch
  concourse: create, delete, get, list, patch, update, watch

ApiGroup: webhook.cert-manager.io
mutations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
validations:
  non-prod:  get, list, watch
  prod:      get, list, watch
  concourse: get, list, watch
```
