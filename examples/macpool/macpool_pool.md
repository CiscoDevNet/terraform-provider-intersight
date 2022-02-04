### Resource Creation
For a pool of 1000 MAC addresses.
```hcl
resource "intersight_macpool_pool" "cisco_af_1" {
  name = "cisco_af_1"

  mac_blocks {
    from = "00:25:B5:AF:10:00"
    size = 1000
  }
}
```
For a pool of 16000 MAC addresses.
```hcl
resource "intersight_macpool_pool" "cisco_0F" {
  name = "cisco_0f"
  tags = [local.terraform]
  dynamic "mac_blocks" {
    for_each = formatlist("%X", range(0, 16))
    content {
      from = "00:25:B5:0F:${mac_blocks.value}0:00"
      size = 1000
    }
  }
}
```
