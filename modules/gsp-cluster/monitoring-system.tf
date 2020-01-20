data "aws_iam_policy_document" "cloudwatch_log_shipping_policy" {
  statement {
    effect = "Allow"

    actions = [
      "logs:DescribeLogGroups",
    ]

    resources = ["*"]
  }

  statement {
    effect = "Allow"

    actions = [
      "logs:DescribeLogStreams",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      aws_cloudwatch_log_group.logs.arn,
      aws_cloudwatch_log_group.application_logs.arn,
      aws_cloudwatch_log_group.dataplane_logs.arn,
      aws_cloudwatch_log_group.host_logs.arn,
    ]
  }
}

data "aws_iam_policy_document" "trust_fluentd_cloudwatch" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type        = "Federated"
      identifiers = [module.k8s-cluster.oidc_provider_arn]
    }

    condition {
      test = "StringEquals"
      variable = "${replace(module.k8s-cluster.oidc_provider_url, "https://", "")}:sub"
      values = ["system:serviceaccount:gsp-system:gsp-fluentd-cloudwatch"]
    }
  }
}

resource "aws_iam_role" "cloudwatch_log_shipping_role" {
  name = "${var.cluster_name}_cloudwatch_log_shipping_role"

  assume_role_policy = data.aws_iam_policy_document.trust_fluentd_cloudwatch.json
}

resource "aws_iam_policy" "cloudwatch_log_shipping_policy" {
  name        = "${var.cluster_name}_cloudwatch_log_shipping_policy"
  description = "Send logs to Clouwatch"

  policy = data.aws_iam_policy_document.cloudwatch_log_shipping_policy.json
}

resource "aws_iam_policy_attachment" "cloudwatch_log_shipping_policy" {
  name = "${var.cluster_name}_cloudwatch_log_shipping_role_policy_attachement"
  roles = [aws_iam_role.cloudwatch_log_shipping_role.name]
  policy_arn = aws_iam_policy.cloudwatch_log_shipping_policy.arn
}

resource "aws_cloudwatch_log_group" "logs" {
  name              = var.cluster_domain
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "application_logs" {
  name              = "/aws/containerinsights/${var.cluster_name}/application"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "host_logs" {
  name              = "/aws/containerinsights/${var.cluster_name}/host"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "dataplane_logs" {
  name              = "/aws/containerinsights/${var.cluster_name}/dataplane"
  retention_in_days = 30
}
