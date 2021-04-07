### Resource Creation

```hcl
resource "intersight_fabric_eth_network_control_policy" "fabric_eth_network_control_policy1" {
    name    = "NCP"
    description = "ethNwControlDes"
    tags    = [
        {
            key = "policy" 
            value   = "ethNwControlPOLICY"
        }
    ]
    cdp_enabled = false
    forgeMac    = "allow"
    lldp_settings   = {
        objectType  = "fabric.LldpSettings"
        receive_enabled = false
        transmit_enabled    = false
    }
    mac_registration_mode = "nativeVlanOnly"
    uplink_fail_action  = "linkDown"
}
```