#!/bin/bash

set -eu -o pipefail

: "${CLUSTER_CONFIG:?}"

CLUSTER_NAME=$(yq -r '.["cluster-name"]' < ${CLUSTER_CONFIG})
PIPELINE_NAME=$(yq -r '.["concourse-pipeline-name"]' < ${CLUSTER_CONFIG})

echo "generating approvers for ${CLUSTER_NAME}..."


fly -t cd-gsp sync

fly -t cd-gsp set-pipeline -p "${PIPELINE_NAME}" \
	--config "pipelines/deployer/deployer.yaml" \
	--load-vars-from "pipelines/deployer/deployer.defaults.yaml" \
	--load-vars-from "${CLUSTER_CONFIG}" \
	--yaml-var 'config-approvers=[]' \
	--yaml-var 'trusted-developer-keys=[]' \
	--check-creds "$@"

fly -t cd-gsp expose-pipeline -p "${PIPELINE_NAME}"

