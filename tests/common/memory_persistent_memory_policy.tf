
resource "intersight_memory_persistent_memory_policy" "memory_persistent_memory_policy1" {
  name        = "memory_persistent_memory_policy1"
  description = "memory persistent memory policies"
  goals = [{
    object_type            = "memory.PersistentMemoryGoal"
    persistent_memory_type = "app-direct-non-interleaved"
    socket_id              = "All Sockets"
    additional_properties  = ""
    class_id               = "memory.PersistentMemoryGoal"
    memory_mode_percentage = 50
  }]
  local_security {
    object_type       = "memory.PersistentMemoryLocalSecurity"
    enabled           = true
    secure_passphrase = "ChangeMe"
  }
  logical_namespaces = [
    {
      additional_properties = ""
      class_id              = "memory.PersistentMemoryLocalSecurity"
      name                  = "logical_namespace_test"
      capacity              = 131072
      object_type           = "memory.PersistentMemoryLocalSecurity"
      mode                  = "block"
      socket_id             = 1
      socket_memory_id      = 6
    }
  ]
  management_mode   = "configured-from-intersight"
  retain_namespaces = false
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
