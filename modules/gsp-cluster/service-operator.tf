resource "aws_s3_bucket" "service-operator" {
  bucket = "gsp-service-operator-${var.cluster_name}"
  acl    = "private"

  tags = {
    Name = "Bucket to store CloudFormation templates generated by the GSP Service Operator"
  }
}

data "aws_iam_policy_document" "trust_svcop" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type        = "Federated"
      identifiers = [module.k8s-cluster.oidc_provider_arn]
    }

    condition {
      test     = "StringEquals"
      variable = "${replace(module.k8s-cluster.oidc_provider_url, "https://", "")}:sub"
      values   = ["system:serviceaccount:gsp-system:gsp-service-operator-service-account"]
    }
  }
}

resource "aws_iam_role" "gsp-service-operator" {
  name               = "${var.cluster_name}-service-operator"
  description        = "Role the service operator assumes"
  assume_role_policy = data.aws_iam_policy_document.trust_svcop.json
}

data "aws_iam_policy_document" "service-operator" {
  statement {
    actions = [
      "cloudformation:*",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    actions = [
      "ec2:DescribeAccountAttributes",
      "ec2:DescribeSecurityGroups",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    actions = [
      "secretsmanager:*",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    actions = [
      "s3:GetObject",
      "s3:PutObject",
    ]

    resources = [
      aws_s3_bucket.service-operator.arn,
      "${aws_s3_bucket.service-operator.arn}/*",
    ]
  }

  statement {
    actions = [
      "rds:*",
      "sqs:*",
      "s3:*",
      "ecr:*",
      "elasticache:*",
    ]

    resources = [
      "*",
    ] # TODO: revisit this
  }

  statement {
    actions = [
      "iam:AttachRolePolicy",
      "iam:CreatePolicy",
      "iam:CreateRole",
      "iam:DeletePolicy",
      "iam:DetachRolePolicy",
      "iam:PutRolePolicy",
      "iam:TagRole",
      "iam:UntagRole",
    ]

    resources = [
      "*",
    ]

    condition {
      test     = "StringEquals"
      variable = "iam:PermissionsBoundary"
      values   = [aws_iam_policy.service-operator-managed-role-permissions-boundary.arn]
    }
  }

  # No iam:PermissionsBoundary context key set on GetRole, DeleteRole
  statement {
    actions = [
      "iam:GetRole",
      "iam:DeleteRole",
      "iam:DeleteRolePolicy",
      "iam:GetRolePolicy",
      "iam:UpdateAssumeRolePolicy",
      "iam:UpdateRole",
    ]

    resources = [
      "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/svcop-${var.cluster_name}-*",
    ]
  }

  statement {
    actions = [
      "iam:CreateServiceLinkedRole",
    ]

    resources = [
      "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/aws-service-role/rds.amazonaws.com/AWSServiceRoleForRDS",
    ]

    condition {
      test     = "StringLike"
      variable = "iam:AWSServiceName"

      values = [
        "rds.amazonaws.com",
      ]
    }
  }
}

resource "aws_iam_policy" "service-operator" {
  name        = "${var.cluster_name}-service-operator"
  description = "Policy for the service operator"
  policy      = data.aws_iam_policy_document.service-operator.json
}

resource "aws_iam_policy_attachment" "service-operator" {
  name       = "${var.cluster_name}-service-operator"
  roles      = [aws_iam_role.gsp-service-operator.name]
  policy_arn = aws_iam_policy.service-operator.arn
}

data "aws_iam_policy_document" "service-operator-managed-role-permissions-boundary" {
  statement {
    actions = [
      "ecr:*",
      "elasticache:*",
      "rds-data:*",
      "rds:*",
      "s3:DeleteObject",
      "s3:DeleteObjectVersion",
      "s3:GetAccelerateConfiguration",
      "s3:GetAnalyticsConfiguration",
      "s3:GetBucketAcl",
      "s3:GetBucketCORS",
      "s3:GetBucketLocation",
      "s3:GetBucketLogging",
      "s3:GetBucketNotification",
      "s3:GetBucketObjectLockConfiguration",
      "s3:GetBucketPolicy",
      "s3:GetBucketPolicyStatus",
      "s3:GetBucketPublicAccessBlock",
      "s3:GetBucketRequestPayment",
      "s3:GetBucketTagging",
      "s3:GetBucketVersioning",
      "s3:GetBucketWebsite",
      "s3:GetEncryptionConfiguration",
      "s3:GetInventoryConfiguration",
      "s3:GetLifecycleConfiguration",
      "s3:GetMetricsConfiguration",
      "s3:GetObject",
      "s3:GetObjectAcl",
      "s3:GetObjectLegalHold",
      "s3:GetObjectRetention",
      "s3:GetObjectTagging",
      "s3:GetObjectTorrent",
      "s3:GetObjectVersion",
      "s3:GetObjectVersionAcl",
      "s3:GetObjectVersionForReplication",
      "s3:GetObjectVersionTagging",
      "s3:GetObjectVersionTorrent",
      "s3:GetReplicationConfiguration",
      "s3:ListBucket",
      "s3:ListBucketByTags",
      "s3:ListBucketMultipartUploads",
      "s3:ListBucketVersions",
      "s3:ListMultipartUploadParts",
      "s3:PutBucketObjectLockConfiguration",
      "s3:PutObject",
      "s3:PutObjectLegalHold",
      "s3:PutObjectRetention",
      "s3:PutObjectVersionAcl",
      "s3:ReplicateObject",
      "s3:RestoreObject",
      "sqs:SendMessage",
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes",
    ]

    resources = [
      "*",
    ]
  }
}

resource "aws_iam_policy" "service-operator-managed-role-permissions-boundary" {
  name        = "${var.cluster_name}-service-operator-managed-role-permissions-boundary"
  description = "Permissions boundary for roles created by the service operator"
  policy      = data.aws_iam_policy_document.service-operator-managed-role-permissions-boundary.json
}

resource "aws_security_group" "rds-from-worker" {
  name        = "${var.cluster_name}_rds_from_worker"
  description = "Allow SQL traffic from worker nodes to RDS instances"
  vpc_id      = var.vpc_id

  ingress {
    from_port = 3306
    to_port   = 3306
    protocol  = "tcp"
    security_groups = [
      module.k8s-cluster.worker_security_group_id,
      module.k8s-cluster.ci_security_group_id,
    ]
  }

  ingress {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"
    security_groups = [
      module.k8s-cluster.worker_security_group_id,
      module.k8s-cluster.ci_security_group_id,
    ]
  }
}

resource "aws_db_subnet_group" "private" {
  name       = "${var.cluster_name}-private"
  subnet_ids = var.private_subnet_ids
}

resource "aws_security_group" "redis-from-worker" {
  name        = "${var.cluster_name}_redis_from_worker"
  description = "Allow Redis traffic from worker nodes to Redis instances"
  vpc_id      = var.vpc_id

  ingress {
    from_port = 6379
    to_port   = 6379
    protocol  = "tcp"
    security_groups = [
      module.k8s-cluster.worker_security_group_id,
      module.k8s-cluster.ci_security_group_id,
    ]
  }
}

resource "aws_elasticache_subnet_group" "private" {
  name       = "${var.cluster_name}-private"
  subnet_ids = var.private_subnet_ids
}
