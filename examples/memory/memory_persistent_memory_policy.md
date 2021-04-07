### Resource Creation

```hcl
resource "intersight_memory_persistent_memory_policy" "memory_persistent_memory_policy1" {
    name            = "memory_persistent_memory_policy1"
    description     = "memory persistent memory policies"
    goals {
        object_type     = "memory.PersistentMemoryGoal"
        persistent_memory_type  = "app-direct"
        socket_id       = "All Sockets"
    }
    local_security {
        object_type     = "memory.PersistentMemoryLocalSecurity"
        enabled         = true
        secure_passphrase   = "ChangeMe"
    }
    logical_name_spaces {
        {
            object_type     = "memory.PersistentMemoryLocalSecurity"
            mode            = "block"
            socket_id       = 1
            socket_memory_id    = 6
        }
    }
    management_mode     = "configured-from-intersight"
    retain_name_spaces  = false
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```