### Resource Creation

```hcl
resource "intersight_fabric_uplink_pc_role" "fabric_uplink_pc_role1" {
  ports = [
    {
      port_id               = 1
      aggregate_port_id     = 0
      slot_id               = 1
      class_id              = "fabric.PortIdentifier"
      object_type           = "fabric.PortIdentifier"
      additional_properties = ""
    },
    {
      port_id               = 2
      aggregate_port_id     = 0
      slot_id               = 1
      class_id              = "fabric.PortIdentifier"
      object_type           = "fabric.PortIdentifier"
      additional_properties = ""
    }
  ]
  admin_speed = "Auto"
  port_policy {
    moid = var.fabric_port_policy1
  }
}

variable "fabric_port_policy1" {
  type        = string
  description = "Fabric port policy Moid"
}
```