### Resource Creation

```hcl
resource "intersight_kubernetes_sys_config_policy" "kubernetes_sys_config_policy1" {
    description = "kubernetes sys config policy"
    name = "kubernetes_sys_config_policy1"
    dns_servers = ["8.8.8.8"]
    ntp_servers = ["10.10.50.50"]
    timezone = "Asia/Calcutta"
    dns_domain_name = "cisco.com"
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```