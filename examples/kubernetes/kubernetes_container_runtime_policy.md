### Resource Creation

```hcl
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
  docker_no_Proxy = [
    "*.test.example.com",
    ".example2.com",
    "127.0.0.0/8"
  ]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```