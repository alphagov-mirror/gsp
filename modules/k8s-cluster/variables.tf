variable "vpc_id" {
  type = string
}

variable "public_subnet_ids" {
  type = list(string)
}

variable "private_subnet_ids" {
  type = list(string)
}

variable "cluster_name" {
  type = string
}

variable "apiserver_allowed_cidrs" {
  type = list(string)
}

variable "eks_version" {
  type = string
}

variable "worker_eks_version" {
  type = string
}

variable "worker_generation_timestamp" {
  type    = string
  default = "none"
}

variable "minimum_workers_per_az_count" {
  type    = string
  default = "1"
}

variable "desired_workers_per_az_map" {
  type    = map(number)
  default = {}
}

variable "maximum_workers_per_az_count" {
  type    = string
  default = "5"
}

variable "worker_on_demand_base_capacity" {
  type    = "string"
  default = "1"
}

variable "worker_on_demand_percentage_above_base" {
  type    = "string"
  default = "100"
}
