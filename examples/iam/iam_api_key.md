### Resource Creation

```hcl
resource "intersight_iam_api_key" "iam_api_key1" {
  hash_algorithm = "SHA256"
  key_spec = [{
    additional_properties = ""
    class_id              = "pkix.RsaAlgorithm"
    modulus               = 2048
    name                  = "RSA"
    object_type           = "pkix.RsaAlgorithm"
  }]
   parent = {
     moid        = var.iam_user
     object_type = "iam.User"
   }
   permission = {
     moid        = var.iam_permission
     object_type = "iam.Permission"
   }
   purpose           = "admin api"
   shared_scope      = ""
   signing_algorithm = "RSASSA-PKCS1-v1_5"
   user = {
     moid        = var.iam_user
     object_type = "iam.User"
   }
}

 variable "iam_permission" {
   type = string
   description = "value for iam_permission"
 }

 variable "iam_user" {
   type = string
   description = "value for iam_user"
 }

 variable "iam_role" {
   type = string
   description = "value for iam_role"
 }
```