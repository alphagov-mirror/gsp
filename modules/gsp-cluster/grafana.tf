resource "aws_iam_role" "grafana" {
  name               = "${var.cluster_name}-grafana"
  description        = "Role the Grafana process assumes"
  assume_role_policy = "${data.aws_iam_policy_document.trust_kiam_server.json}"
}

data "aws_iam_policy_document" "grafana_cloudwatch" {
  statement {
    effect = "Allow"

    actions = [
      "cloudwatch:ListMetrics",
      "cloudwatch:GetMetricStatistics",
      "cloudwatch:GetMetricData",
    ]

    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "grafana" {
  name   = "${var.cluster_name}-grafana"
  role   = "${aws_iam_role.grafana.id}"
  policy = "${data.aws_iam_policy_document.grafana_cloudwatch.json}"
}
