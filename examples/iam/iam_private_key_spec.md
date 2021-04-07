### Resource Creation

```hcl
resource "intersight_iam_private_key_spec" "iam_private_key_spec1" {
    certificate_request {
        email_address = "email@example.com"
        name = "iam_certificate_request1"
        self_signed = false
    }
}
```