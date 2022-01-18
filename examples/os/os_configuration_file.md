### Resource Creation

```hcl
resource "intersight_os_configuration_file" "os_configuration_file1" {
  name = "os_configuration_file1"
   placeholders = [{
     additional_properties = ""
     class_id ="os.PlaceHolder"
     value = ""
     object_type  = "os.PlaceHolder"
     is_value_set = true
     type = [{
       object_type = "workflow.PrimitiveDataType"
       additional_properties = ""
       class_id = "workflow.PrimitiveDataType"
       input_parameters = ""
       label = ""
       name = ""
       properties = [{
         additional_properties = ""
         class_id = "workflow.PrimitiveDataType"
         constraints = [""]
         inventory_selector = ""
         object_type = "workflow.PrimitiveDataType"
         secure = false
         type = ""
       }]
       required = ""
       default = {
         object_type = "workflow.DefaultValue"
         override    = false
         value       = "<default value of the data_type>"
       }
       description = "Description for the Default data type"
       display_meta = [{
         additional_properties = ""
         class_id = "workflow.DisplayMeta"
         object_type        = "workflow.DisplayMeta"
         inventory_selector = true
         widget_type        = "Radio"
       }]
     }]
  }]
   catalog {
     object_type = "os.Catalog"
     moid        = var.catalog
   }
   version {
     object_type = "hcl.OperatingSystem "
     moid        = var.hcl_OperatingSystem
   }
}

 variable "catalog" {
   type = string
   description = "value for catalog"
 }

variable "hcl_OperatingSystem" {
   type = string
   description = "value for hcl_OperatingSystem"
 }
```