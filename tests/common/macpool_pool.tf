resource "intersight_macpool_pool" "AUTO_test_macpool1" {
  name = "AUTO_test_macpool1"

  mac_blocks {
    from = "00:25:B5:AF:10:00"
    size = 100
  }
}

resource "intersight_macpool_pool" "AUTO_test_macpool2" {
  name = "AUTO_test_macpool2"
    tags {
    key   = "testing"
    value = "macPool"
  }
  dynamic "mac_blocks" {
    for_each = formatlist("%X", range(0, 16))
    content {
      from = "00:25:B5:0F:${mac_blocks.value}0:00"
      size = 10
    }
  }
}


