# ADR045: Dev namespaces

## Status

Pending

## Context

Following the retirement of gsp-local the service teams have no way of testing
changes before a merge to master. So the feedback cycle is very slow, and
potentially risky. We need to enable devs to get more rapid feedback through use
of namespaces that allow for deployments from arbitrary branches or from local
`kubectl apply` actions.

This presents several security problems that need to be addressed.

The in-cluster concourse:

* must not be able to create or edit daemonsets
* must not be able to create pods that use host networking, or run in privildged
  mode
* must not be authorised to create any "cluster" scoped resources (e.g.
  clusterrole, custom resources etc.)

For harbor:

* a dev in one namespace must not be authorised to push, edit or delete images
  relating to namespaces outside their own
* notary signing keys must be namespace-scoped

For external-dns:

* a namespace must not be able to hijack the DNS entries of another namespace

For istio:

* any istio resources deployed as part of a namespace must not have any impact
  on other namespaces

For CloudHSM access:

* connectivity only enabled via `-cluster-config`, as with "prod" namespaces and
  pods
* credentials should be different for each namespace, and not the same as those
  used for the "production" instances of the application

## Decision

We will address the majority of the security concerns by implementing [ADR043].

We will address the harbor concerns by creating namespace-scoped credentials
relating to namespace-scoped harbor "projects" and provide these credentials via
secrets in the namespace.

We will address the DNS concerns by locking down each namespace instance of
external-dns to a dedicated zone.

We will address the istio concerns through the use of gatekeeper constraints
(e.g. all istio resources that support it have `exportTo: ["."]` set).

We will address the CloudHSM concerns by creating users (as required) in the
CloudHSM cluster and providing those credentials into the namespace via secrets
(or SealedSecrets).

## Consequences

Users will be able to arbitrarily `kubectl apply` (or equivalent) resources into
their own namespaces without affecting others, the platform infrastructure or
production systems. It will be possible to configure concourse to build and
deploy changes from a branch within a Team (scoped to a namespace) allowing more
rapid feedback of proposed changes. The potential impact of security attacks or
compromises will be limited by namespace-isolated credentials.

[ADR043]: ./ADR043-k8s-resource-access.md