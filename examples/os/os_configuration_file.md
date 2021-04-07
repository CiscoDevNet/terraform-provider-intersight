### Resource Creation

```hcl
resource "intersight_os_configuration_file" "os_configuration_file1" {
    name = "os_configuration_file1"
    placeholders = [{
        object_type = "os.PlaceHolder"
        is_value_set = true
        type {
            object_type = "workflow.PrimitiveDataType"
            default {
                object_type = "workflow.DefaultValue"
                override = false
                value = "<default value of the data_type>"
            }
            description = "Description for the Default data type"
            display_meta {
                object_type = "workflow.DisplayMeta"
                inventory_selector = true
                widget_type = "Radio"
            }
        }
    }]
    catalog {
        object_type = "os.Catalog"
        moid = var.catalog
    }
    version {
        object_type  = "hcl.OperatingSystem "
        moid = var.hcl.OperatingSystem
    }
}
```