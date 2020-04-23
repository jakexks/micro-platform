resource "kubernetes_namespace" "network_namespace" {
  metadata {
    name = var.network_namespace
  }
}



// One day this can just be a for_each ["list", "of", "services"]
// https://github.com/hashicorp/terraform/issues/10462#issuecomment-575738220
module "api" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  micro_image = var.micro_image

  service_name = "api"
  service_port = 443

  extra_env_vars = {
    "MICRO_API_NAMESPACE"  = "domain"
    "MICRO_ENABLE_STATS"   = "true"
    "MICRO_ENABLE_ACME"    = "true"
    "MICRO_ACME_PROVIDER"  = "certmagic"
    "MICRO_ACME_HOSTS"     = "*.${var.domain_name},*.cloud.${var.domain_name},${var.domain_name}"
    "MICRO_STORE"          = "service"
    "MICRO_STORE_DATABASE" = "micro"
    "MICRO_STORE_TABLE"    = "micro"
    "CF_API_TOKEN"         = var.cloudflare_api_token
  }
}

module "api_auth" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  micro_image = var.micro_image

  service_name = "api.auth"
  service_port = 8000

  extra_env_vars = {
    "MICRO_AUTH" = "service"
  }
}

module "auth" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  micro_image = var.micro_image

  service_name = "auth"
  service_port = 8010

  extra_env_vars = {
    "MICRO_AUTH" = "service"
  }
}


module "broker" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  micro_image = var.micro_image

  service_name = "broker"
  service_port = 8001
}

module "config" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  micro_image = var.micro_image

  service_name = "config"
  service_port = 8080
}

module "debug_web" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name       = "debug-web"
  create_k8s_service = false

  extra_env_vars = {
    "MICRO_NETDATA_URL" = "http://netdata.${var.resource_namespace}.svc:19999"
  }
}

module "debug" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name       = "debug"
  create_k8s_service = false
}

module "network_api" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name       = "network-api"
  create_k8s_service = false

  extra_env_vars = {
    "MICRO_SERVER_ADDRESS" = "0.0.0.0:9090"
  }
}

module "proxy" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name = "proxy"
  service_port = 8081
}

module "registry" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name = "registry"
  service_port = 8000
}

module "router" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name = "router"
  service_port = 8084
}

module "store" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name       = "store"
  create_k8s_service = false

  extra_env_vars = {
    "MICRO_STORE_BACKEND" = "cockroach"
    "MICRO_STORE_NODES"   = "host=cockroachdb-public.${var.resource_namespace}.svc port=26257 sslmode=disable user=root"
  }
}

module "web" {
  source = "./service"

  resource_namespace = var.resource_namespace
  network_namespace  = var.network_namespace

  service_name = "web"
  service_port = 443

  extra_env_vars = {
    "MICRO_ENABLE_ACME"   = "true"
    "MICRO_ACME_PROVIDER" = "certmagic"
    "MICRO_ACME_HOSTS"    = "*.${var.domain_name},*.cloud.${var.domain_name},${var.domain_name}"
  }
}
