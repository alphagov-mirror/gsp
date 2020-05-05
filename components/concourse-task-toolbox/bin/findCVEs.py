#!/usr/bin/python3

import subprocess
import json
import sys

from kubernetes import client, config
# whitelists against images that are problematic to pull/scan
GLOBAL_IMAGE_WHITELIST = [
    'jaegertracing/all-in-one:1.16', # error in image scan: scan failed: failed to apply layers: unknown OS - no shell, no ls - possibly scratch
    'k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1', # error in image scan: scan failed: failed to apply layers: unknown OS - no shell, no ls - possibly scratch
    'k8s.gcr.io/metrics-server-amd64:v0.3.0', # error in image scan: scan failed: failed to apply layers: unknown OS
    'quay.io/calico/typha:v3.8.1', # error in image scan: scan failed: failed to apply layers: unknown OS
    'quay.io/coreos/configmap-reload:v0.0.1', # error in image scan: scan failed: failed to apply layers: unknown OS
    'quay.io/coreos/prometheus-config-reloader:v0.37.0', # error in image scan: scan failed: failed to apply layers: unknown OS
    'quay.io/coreos/prometheus-operator:v0.37.0', # error in image scan: scan failed: failed to apply layers: unknown OS
    'quay.io/prometheus/node-exporter:v0.18.1', # error in image scan: scan failed: failed to apply layers: unknown OS
    'quay.io/prometheus/prometheus:v2.15.2', # error in image scan: scan failed: failed to apply layers: unknown OS
]
GLOBAL_IMAGE_SOURCE_WHITELIST = [
    '.dkr.ecr.eu-west-2.amazonaws.com/', # ECR
    '.dkr.ecr.us-west-2.amazonaws.com/', # ECR - for EKS upstream
]

# This is to work around trivy being unwilling to work with their seemingly broken responses, see https://github.com/aquasecurity/trivy/issues/401#issuecomment-611454832
PULL_FIRST = [
    'quay.io/bitnami/sealed-secrets-controller:v0.7.0',
    'quay.io/calico/node:v3.8.1',
    'quay.io/open-policy-agent/gatekeeper:v3.0.4-beta.1',
    'quay.io/kiali/kiali:v1.9',
]

# whitelists against vulnerabilities we've considered for various reasons


def whitelisted(vulnerability):
    if vulnerability['image_name'].startswith('fluent/fluentd-kubernetes-daemonset:v1.') and \
       vulnerability['vulnerability']['VulnerabilityID'] == 'CVE-2020-8130':
        # this shows up in usr/local/bundle/gems/async-http-0.50.0/examples/fetch/Gemfile.lock -
        # which is just an example in one of the libraries, and also in
        # usr/local/bundle/gems/http_parser.rb-0.6.0/Gemfile.lock
        # The second one is slightly more concerning but the nature of the vulnerability appears
        # to be unwanted behaviour from some internal functions of a build library, which seems
        # unlikely to pose a real problem for us.
        # In https://hackerone.com/reports/651518 it was written:
        # "the attack surface was limited because if It's difficult to inject malicious input to
        # Rake::FileList by attackers with the current usage of Rake in the world."
        return True
    return False

trivy_cache = {}
config.load_kube_config()
vulnerabilities = []
exceptions_encountered = False

for pod in client.CoreV1Api().list_pod_for_all_namespaces(watch=False).items:
    for container in pod.spec.containers:
        image_name = container.image.replace('docker.io/', '')
        if image_name in GLOBAL_IMAGE_WHITELIST:
            continue
        if any(source in image_name for source in GLOBAL_IMAGE_SOURCE_WHITELIST):
            continue
        if image_name not in trivy_cache:
            trivy_cache[image_name] = []
            if image_name in PULL_FIRST:
                subprocess.check_call(['/usr/bin/docker', 'pull', image_name])
            try:
                output = subprocess.check_output([
                    'trivy',
                    '--format', 'json',
                    '--quiet',
                    '--ignore-unfixed', # remove this if you want to learn about CVE-2005-2541
                    '-s', 'CRITICAL',
                    image_name
                ])
            except subprocess.CalledProcessError as e:
                print(e)
                exceptions_encountered = True
                continue
            for target in json.loads(output):
                trivy_cache[image_name] += target.get('Vulnerabilities') or []
        for trivy_vulnerability_obj in trivy_cache[image_name]:
            vulnerability = {
                'namespace': pod.metadata.namespace,
                'container_name': container.name,
                'image_name': image_name,
                'vulnerability': trivy_vulnerability_obj,
            }
            # de-duplicate multiple pods belonging to the same ReplicaSet/StatefulSet/DaemonSet etc. by attributing to their owning object
            if len(pod.metadata.owner_references) > 0:
                assert len(pod.metadata.owner_references) == 1
                vulnerability['kind'] = pod.metadata.owner_references[0].kind
                vulnerability['name'] = pod.metadata.owner_references[0].name
            else:
                vulnerability['kind'] = 'Pod'
                vulnerability['name'] = pod.metadata.name
            if whitelisted(vulnerability):
                continue
            if vulnerability not in vulnerabilities:
                vulnerabilities.append(vulnerability)
                print(json.dumps(vulnerability, indent=4))

if len(vulnerabilities) > 0 or exceptions_encountered:
    sys.exit(1)
