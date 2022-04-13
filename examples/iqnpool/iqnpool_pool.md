### Resource Creation

```hcl
resource "intersight_iqnpool_pool" "iqnpool_pool1" {
  name             = "ippool_pool1"
  description      = "ippool pool"
  assignment_order = "sequential"
  prefix           = "iqn1.alpha.com"
  iqn_suffix_blocks {
    object_type = "iqn.SuffixBlocks"
    suffix      = "alphadc-1"
  }
}
```