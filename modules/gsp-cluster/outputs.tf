output "kubeconfig" {
  value = module.k8s-cluster.kubeconfig
}

output "worker_security_group_id" {
  value = module.k8s-cluster.worker_security_group_id
}

output "values" {
  sensitive = true
  value     = data.template_file.values.rendered
}

output "trust_kiam_server_policy_json" {
  value = data.aws_iam_policy_document.trust_kiam_server.json
}
