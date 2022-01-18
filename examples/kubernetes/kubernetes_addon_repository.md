### Resource Creation

```hcl
resource "intersight_kubernetes_addon_repository" "kubernetes_addon_repository1" {
  insecure_skip_verification = false
  name                       = "kubernetes_addon_repository1"
  username                   = "admin1"
}
```