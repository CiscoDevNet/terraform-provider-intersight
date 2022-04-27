### Resource Creation

```hcl
resource "intersight_fabric_uplink_pc_role" "fabric_uplink_pc_role1" {
  ports {
    port_id           = 1
    aggregate_port_id = 0
    slot_id           = 1
    class_id          = "fabric.PortIdentifier"
    object_type       = "fabric.PortIdentifier"
  }
  ports {
    port_id           = 2
    aggregate_port_id = 0
    slot_id           = 1
    class_id          = "fabric.PortIdentifier"
    object_type       = "fabric.PortIdentifier"
  }
  admin_speed = "Auto"
  port_policy {
    moid        = var.fabric_port_policy
    object_type = "fabric.PortPolicy"
  }
}

variable "fabric_port_policy" {
  type        = string
  description = "Fabric port policy Moid"
}
```