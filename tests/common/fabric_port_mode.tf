
resource "intersight_fabric_port_mode" "fabric_port_mode1" {
  custom_mode   = "FibreChannel"
  port_id_end   = 8
  port_id_start = 1
  slot_id       = 1
  port_policy {
    moid = intersight_fabric_port_policy.fabric_port_policy1.moid
  }
}
