### Resource Creation

```hcl
resource "intersight_sdwan_profile" "sdwan_profile1" {
  name        = "sdwan_profile1"
  description = "sdwan_profile"
  src_template {
    object_type = "policy.abstractprofile"
    moid        = var.policy_abstractprofile
  }
  config_context {
    object_type    = "policy.ConfigContext"
    control_action = "deploy"
  }
  router_nodes = [{
    object_type = "sdwan.RouterNode"
    moid        = var.sdwan_router_node
    class_id = "sdwan.RouterNode"
    additional_properties = ""
    selector = ""
  }]
  router_policy {
    object_type = "router.policy"
    moid        = var.router_policy
  }
  vmanage_account {
    object_type = "sdwan.VmanageAccountPolicy"
    moid        = var.sdwan_vmanage_account_policy
  }
}

variable "policy_abstractprofile" {
  type = string
  description = "value for policy_abstractprofile"
}

variable "sdwan_router_node" {
  type = string
  description = "value for sdwan_router_node"  
}

variable "router_policy" {
  type = string
  description = "value for router policy"
}

variable "sdwan_vmanage_account_policy" {
  type = string
  description = "value for sdwan vmange account policy"
}
```