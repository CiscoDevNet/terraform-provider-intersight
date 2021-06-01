### Resource Creation

```hcl
resource "intersight_macpool_pool" "macpool_pool1" {
  name = "AUTO_test_macpool"
  mac_blocks {
    object_type = "macpool.Block"
    from        = "00:25:B5:9d:68:16"
    to          = "70:DF:2F:87:06:40"
  }
}
```