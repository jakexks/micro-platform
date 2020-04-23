variable "network_namespace" {
  type = string
}

variable "resource_namespace" {
  type = string
}

variable "domain_name" {
  type = string
}

variable "micro_image" {
  type    = string
  default = "micro/micro:latest"
}

variable "cloudflare_api_token" {
  type = string
}