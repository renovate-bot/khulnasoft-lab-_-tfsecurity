#tfsecurity:ignore:AWS002 tfsecurity:ignore:AWS098
resource "aws_s3_bucket" "web_bucket" {
  bucket   = "randomness-123123213"
  acl      = "public-read" #tfsecurity:ignore:AWS001

  website {
    index_document = "index.html"
  }

  versioning {
    enabled = true
  }

  tags = {
    Name      = "randomness-123123213"
    terraform = "true"
  }
}
