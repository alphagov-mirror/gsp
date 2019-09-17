resource "aws_s3_bucket" "service-operator" {
  bucket = "gsp-service-operator-${var.cluster_name}"
  acl    = "private"

  tags = {
    Name = "Bucket to store CloudFormation templates generated by the GSP Service Operator"
  }
}

resource "aws_iam_role" "gsp-service-operator" {
  name               = "${var.cluster_name}-service-operator"
  description        = "Role the service operator assumes"
  assume_role_policy = "${data.aws_iam_policy_document.trust_kiam_server.json}"
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
      "${aws_s3_bucket.service-operator.arn}",
      "${aws_s3_bucket.service-operator.arn}/*",
    ]
  }

  statement {
    actions = [
      "rds:*",
      "sqs:*",
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
      values   = ["${aws_iam_policy.service-operator-managed-role-permissions-boundary.arn}"]
    }
  }

  # No iam:PermissionsBoundary context key set on GetRole, DeleteRole
  statement {
    actions = [
      "iam:GetRole",
      "iam:DeleteRole",
      "iam:DeleteRolePolicy",
      "iam:GetRolePolicy",
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
  policy      = "${data.aws_iam_policy_document.service-operator.json}"
}

resource "aws_iam_policy_attachment" "service-operator" {
  name       = "${var.cluster_name}-service-operator"
  roles      = ["${aws_iam_role.gsp-service-operator.name}"]
  policy_arn = "${aws_iam_policy.service-operator.arn}"
}

data "aws_iam_policy_document" "service-operator-managed-role-permissions-boundary" {
  statement {
    actions = [
      "rds-data:*",
      "sqs:*",
      "rds:*",
    ]

    resources = [
      "*",
    ]
  }
}

resource "aws_iam_policy" "service-operator-managed-role-permissions-boundary" {
  name        = "${var.cluster_name}-service-operator-managed-role-permissions-boundary"
  description = "Permissions boundary for roles created by the service operator"
  policy      = "${data.aws_iam_policy_document.service-operator-managed-role-permissions-boundary.json}"
}

resource "aws_security_group" "rds-from-worker" {
  name        = "rds_from_worker"
  description = "Allow SQL traffic from worker nodes to RDS instances"
  vpc_id      = "${var.vpc_id}"

  ingress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
    security_groups = ["${module.k8s-cluster.worker_security_group_id}"]
  }
}

resource "aws_db_subnet_group" "private" {
  name       = "sandbox-private"
  subnet_ids = ["${var.private_subnet_ids}"]
}
