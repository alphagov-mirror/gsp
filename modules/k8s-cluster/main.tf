data "aws_caller_identity" "current" {
}

data "aws_region" "current" {
}

data "aws_subnet" "private_subnets" {
  count = length(var.private_subnet_ids)
  id    = element(var.private_subnet_ids, count.index)
}

resource "aws_eks_cluster" "eks-cluster" {
  name     = var.cluster_name
  role_arn = aws_iam_role.eks-cluster.arn
  version  = var.eks_version

  vpc_config {
    security_group_ids = [aws_security_group.controller.id]
    subnet_ids         = concat(var.private_subnet_ids, var.public_subnet_ids)
  }

  enabled_cluster_log_types = [
    "api",
    "audit",
    "authenticator",
    "controllerManager",
    "scheduler",
  ]

  depends_on = [
    aws_iam_role_policy_attachment.eks-cluster-policy,
    aws_iam_role_policy_attachment.eks-service-policy,
    aws_cloudwatch_log_group.eks,
  ]

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_iam_openid_connect_provider" "eks" {
  client_id_list = ["sts.amazonaws.com"]
  # the thumbprint identifies the root CA for the OIDC provider. See
  # https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html
  # in this case, this thumbprint is the thumbprint of the root CA
  # presented by `oidc.eks.eu-west-2.amazonaws.com`; this isn't likely to change
  # I checked and the same thumbprint applies to eu-west-1
  thumbprint_list = ["9E99A48A9960B14926BB7F3B02E22DA2B0AB7280"]
  url             = aws_eks_cluster.eks-cluster.identity[0].oidc[0].issuer
}

resource "aws_cloudwatch_log_group" "eks" {
  name              = "/aws/eks/${var.cluster_name}/cluster"
  retention_in_days = 30

  lifecycle {
    prevent_destroy = true
  }
}

# As per https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
resource "aws_cloudformation_stack" "worker-nodes" {
  name          = "${var.cluster_name}-worker-nodes"
  template_body = file("${path.module}/data/nodegroup.yaml")
  capabilities  = ["CAPABILITY_IAM"]

  parameters = {
    NodeImageId                         = "/aws/service/eks/optimized-ami/${var.worker_eks_version}/amazon-linux-2/recommended/image_id"
    ClusterName                         = var.cluster_name
    ClusterControlPlaneSecurityGroup    = aws_security_group.controller.id
    NodeGroupName                       = "worker"
    NodeAutoScalingGroupMinSize         = "0"
    NodeAutoScalingGroupDesiredCapacity = "0"
    NodeAutoScalingGroupMaxSize         = "0"
    NodeInstanceType                    = var.worker_instance_type
    NodeVolumeSize                      = "40"
    BootstrapArguments                  = "--kubelet-extra-args \"--node-labels=node-role.kubernetes.io/worker --event-qps=0\""
    VpcId                               = var.vpc_id
    Subnets                             = join(",", var.private_subnet_ids)
  }

  timeouts {
    create = "30m"

    # rolling worker nodes 1 at a time could be time consuming. Stop concourse going red
    update = "90m"
    delete = "30m"
  }

  depends_on = [aws_eks_cluster.eks-cluster]
}

resource "aws_cloudformation_stack" "worker-nodes-per-az" {
  count = length(var.private_subnet_ids)
  name = "${var.cluster_name}-worker-nodes-${element(
    data.aws_subnet.private_subnets.*.availability_zone,
    count.index,
  )}"
  template_body = file("${path.module}/data/nodegroup-v2.yaml")
  capabilities  = ["CAPABILITY_IAM"]

  parameters = {
    NodeImageId                      = "/aws/service/eks/optimized-ami/${var.worker_eks_version}/amazon-linux-2/recommended/image_id"
    ClusterName                      = var.cluster_name
    ClusterControlPlaneSecurityGroup = aws_security_group.controller.id
    NodeGroupName = "worker-${element(
      data.aws_subnet.private_subnets.*.availability_zone,
      count.index,
    )}"
    NodeAutoScalingGroupMinSize = var.minimum_workers_per_az_count
    NodeAutoScalingGroupDesiredCapacity = tostring(lookup(
      var.desired_workers_per_az_map,
      element(
        data.aws_subnet.private_subnets.*.availability_zone,
        count.index,
      ),
      tonumber(var.minimum_workers_per_az_count)
    ))
    NodeAutoScalingGroupMinInstancesInService = tostring(min(
      lookup(
        var.desired_workers_per_az_map,
        element(
          data.aws_subnet.private_subnets.*.availability_zone,
          count.index,
        ),
        tonumber(var.minimum_workers_per_az_count)
      ),
      var.maximum_workers_per_az_count - 1
    ))
    NodeAutoScalingGroupMaxSize = var.maximum_workers_per_az_count

    NodeAutoScalingGroupOnDemandPercentageAboveBase = var.worker_on_demand_percentage_above_base

    NodeInstanceType    = var.worker_instance_type
    NodeInstanceProfile = aws_cloudformation_stack.worker-nodes.outputs["NodeInstanceProfile"]
    NodeVolumeSize      = "40"
    BootstrapArguments  = "--kubelet-extra-args \"--node-labels=node-role.kubernetes.io/worker --event-qps=0\""
    VpcId               = var.vpc_id
    Subnets             = element(data.aws_subnet.private_subnets.*.id, count.index)
    NodeSecurityGroups  = "${aws_security_group.node.id},${aws_security_group.worker.id}"
    NodeTargetGroups    = "${aws_cloudformation_stack.worker-nodes.outputs["HTTPTargetGroup"]},${aws_cloudformation_stack.worker-nodes.outputs["TCPTargetGroup"]}"
  }

  depends_on = [
    aws_eks_cluster.eks-cluster,
    aws_cloudformation_stack.worker-nodes,
  ]
}

resource "aws_autoscaling_lifecycle_hook" "worker-nodes-per-az-lifecycle-hook" {
  count = length(var.private_subnet_ids)
  name = "${var.cluster_name}-worker-${element(
    data.aws_subnet.private_subnets.*.availability_zone,
    count.index,
  )}"
  autoscaling_group_name = lookup(aws_cloudformation_stack.worker-nodes-per-az[count.index].outputs, "AutoScalingGroupName", "")
  default_result         = "ABANDON"
  heartbeat_timeout      = 180
  lifecycle_transition   = "autoscaling:EC2_INSTANCE_TERMINATING"
}

resource "aws_cloudformation_stack" "kiam-server-nodes" {
  name          = "${var.cluster_name}-kiam-server-nodes"
  template_body = file("${path.module}/data/nodegroup.yaml")
  capabilities  = ["CAPABILITY_IAM"]

  parameters = {
    NodeImageId                         = "/aws/service/eks/optimized-ami/${var.worker_eks_version}/amazon-linux-2/recommended/image_id"
    ClusterName                         = var.cluster_name
    ClusterControlPlaneSecurityGroup    = aws_security_group.controller.id
    NodeGroupName                       = "kiam"
    NodeAutoScalingGroupMinSize         = "2"
    NodeAutoScalingGroupDesiredCapacity = "2"
    NodeAutoScalingGroupMaxSize         = "3"
    NodeInstanceType                    = "t3.medium"
    NodeVolumeSize                      = "40"
    BootstrapArguments                  = "--kubelet-extra-args \"--node-labels=node-role.kubernetes.io/cluster-management --register-with-taints=node-role.kubernetes.io/cluster-management=:NoSchedule --event-qps=0\""
    VpcId                               = var.vpc_id
    Subnets                             = join(",", var.private_subnet_ids)
  }

  depends_on = [aws_eks_cluster.eks-cluster]
}

resource "aws_autoscaling_lifecycle_hook" "kiam-nodes-lifecycle-hook" {
  name                   = "${var.cluster_name}-kiam"
  autoscaling_group_name = lookup(aws_cloudformation_stack.kiam-server-nodes.outputs, "AutoScalingGroupName", "")
  default_result         = "ABANDON"
  heartbeat_timeout      = 180
  lifecycle_transition   = "autoscaling:EC2_INSTANCE_TERMINATING"
}

resource "aws_cloudformation_stack" "ci-nodes" {
  name          = "${var.cluster_name}-ci-nodes"
  template_body = file("${path.module}/data/nodegroup.yaml")
  capabilities  = ["CAPABILITY_IAM"]

  parameters = {
    NodeImageId                         = "/aws/service/eks/optimized-ami/${var.worker_eks_version}/amazon-linux-2/recommended/image_id"
    ClusterName                         = var.cluster_name
    ClusterControlPlaneSecurityGroup    = aws_security_group.controller.id
    NodeGroupName                       = "ci"
    NodeAutoScalingGroupMinSize         = var.ci_worker_count
    NodeAutoScalingGroupDesiredCapacity = var.ci_worker_count
    NodeAutoScalingGroupMaxSize         = var.ci_worker_count + 1
    NodeInstanceType                    = var.ci_worker_instance_type
    NodeVolumeSize                      = "75"
    BootstrapArguments                  = "--kubelet-extra-args \"--node-labels=node-role.kubernetes.io/ci --register-with-taints=node-role.kubernetes.io/ci=:NoSchedule --event-qps=0\""
    VpcId                               = var.vpc_id
    Subnets                             = join(",", var.private_subnet_ids)
  }

  depends_on = [aws_eks_cluster.eks-cluster]
}

resource "aws_autoscaling_lifecycle_hook" "ci-nodes-lifecycle-hook" {
  name                   = "${var.cluster_name}-ci"
  autoscaling_group_name = lookup(aws_cloudformation_stack.ci-nodes.outputs, "AutoScalingGroupName", "")
  default_result         = "ABANDON"
  heartbeat_timeout      = 180
  lifecycle_transition   = "autoscaling:EC2_INSTANCE_TERMINATING"
}

data "template_file" "kubeconfig" {
  template = file("${path.module}/data/kubeconfig")

  vars = {
    apiserver_endpoint = aws_eks_cluster.eks-cluster.endpoint
    ca_cert            = aws_eks_cluster.eks-cluster.certificate_authority[0].data
    name               = var.cluster_name
    cluster_id         = var.cluster_name
  }
}

