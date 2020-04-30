locals {
  router_name = "go.micro.router"
  router_port = 8011
  router_labels = merge(
    local.common_labels,
    {
      "name" = local.router_name
    }
  )
  router_annotations = merge(
    local.common_annotations,
    {
      "name" = local.router_name
    }
  )
  router_env = merge(
    local.common_env_vars,
    {
    }
  )
}

module "router_cert" {
  source = "./cert"

  ca_cert_pem        = tls_self_signed_cert.platform_ca_cert.cert_pem
  ca_private_key_pem = tls_private_key.platform_ca_key.private_key_pem
  private_key_alg    = var.private_key_alg

  subject = local.router_name
}

resource "kubernetes_secret" "router_cert" {
  metadata {
    name        = "${replace(local.router_name, ".", "-")}-cert"
    namespace   = var.platform_namespace
    labels      = local.router_labels
    annotations = local.router_annotations
  }
  data = {
    "cert.pem" = module.router_cert.cert_pem
    "key.pem"  = module.router_cert.key_pem
  }
  type = "Opaque"
}

resource "kubernetes_deployment" "router" {
  metadata {
    name        = replace(local.router_name, ".", "-")
    namespace   = var.platform_namespace
    labels      = local.router_labels
    annotations = merge(local.common_annotations, local.router_annotations)
  }
  spec {
    replicas = 1
    selector {
      match_labels = local.router_labels
    }
    template {
      metadata {
        labels = local.router_labels
      }
      spec {
        container {
          name = replace(local.router_name, ".", "-")
          dynamic "env" {
            for_each = local.router_env
            content {
              name  = env.key
              value = env.value
            }
          }
          args              = ["router"]
          image             = var.micro_image
          image_pull_policy = var.image_pull_policy
          port {
            container_port = local.router_port
            name           = "router-port"
          }
          volume_mount {
            mount_path = "/etc/micro/certs"
            name       = "certs"
          }
          volume_mount {
            mount_path = "/etc/micro/ca"
            name       = "platform-ca"
          }
        }
        volume {
          name = "platform-ca"
          secret {
            secret_name  = kubernetes_secret.platform_ca.metadata[0].name
            default_mode = "0600"
            items {
              key  = "ca.pem"
              path = "ca.pem"
            }
          }
        }
        volume {
          name = "certs"
          secret {
            default_mode = "0600"
            secret_name  = kubernetes_secret.router_cert.metadata[0].name
          }
        }
        automount_service_account_token = true
      }
    }
  }
}

resource "kubernetes_service" "router" {
  metadata {
    name        = replace(local.router_name, ".", "-")
    namespace   = var.platform_namespace
    labels      = local.router_labels
    annotations = merge(local.common_annotations, local.router_annotations)
  }
  spec {
    port {
      port        = local.router_port
      target_port = local.router_port
    }
    selector = local.router_labels
  }
}
