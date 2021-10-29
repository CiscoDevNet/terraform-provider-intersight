
resource "intersight_kubernetes_container_runtime_policy" "kubernetes_container_runtime_policy1" {
  description = "kubernetes container runtime policy"
  name        = "kubernetes_container_runtime_policy1"
  docker_http_proxy {
    hostname = "10.1.1.10"
    password = "ChangeMe"
    port     = 3001
    protocol = "http"
    username = "admin1"
  }
  docker_https_proxy {
    hostname = "10.1.1.10"
    password = "ChangeMe"
    port     = 3001
    protocol = "https"
    username = "admin1"
  }
  docker_no_proxy = [""]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
