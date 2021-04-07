### Resource Creation

```hcl
resource "intersight_sdwan_profile" "sdwan_profile1" {
    name = "sdwan_profile1"
    description = "sdwan_profile"
    scr_template {
        object_type = "policy.abstractprofile"
        moid = var.policy_abstractprofile
    }
    config_context {
        object_type = "policy.ConfigContext"
        control_action = "deploy"
    }
    router_nodes = [{
        object_type = "sdwan.RouterNode"
        moid = var.sdwan_router_node
    }]
    router_policy {
        object_type = "router.policy"
        moid = var.router_policy
    }
    vmanage_account {
        object_type = "sdwan.VmanageAccountPolicy"
        moid = var.sdwan_vmanage_account_policy
    }
}
```