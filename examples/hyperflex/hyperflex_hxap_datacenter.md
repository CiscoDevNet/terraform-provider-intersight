### Resource Creation

```hcl
resource "intersight_hyperflex_hxap_datacenter" "hyperflex_hxap_datacenter1" {
  name = "hyperflex_hxap_datacenter1"
  account {
    object_type = "iam.Account"
    moid        = var.iam_account_hxap
  }
}

variable "iam_account_hxap" {
  type        = string
  description = "Moid of Iam account"
}
```