resource "intersight_iam_api_key" "iam_api_key1" {
  hash_algorithm    = "SHA256"
  purpose           = "Test Api Key"
  signing_algorithm = "RSASSA-PKCS1-v1_5"
}

