resource "kubernetes_namespace" "network" {
  metadata {
    name = var.network_namespace
  }
}

output "network_namespace" {
  value = kubernetes_namespace.network.id
}

resource "kubernetes_namespace" "resource" {
  metadata {
    name = var.resource_namespace
  }
}

output "resource_namespace" {
  value = kubernetes_namespace.resource.id
}
