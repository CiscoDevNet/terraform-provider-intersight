### Resource Creation

```hcl
resource "intersight_macpool_pool" "macpool_pool1" {
    name        = "AUTO_test_macpool"
    mac_blocks {
        object_type = "macpool.Block"
        from        = "0025B59d6816"
        to          = "70DF2F870640"
    }
}
```