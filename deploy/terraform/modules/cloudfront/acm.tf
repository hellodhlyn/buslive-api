provider "aws" {
  region = "us-east-1"
  alias  = "us-east-1"
}

data "aws_acm_certificate" "lynlab-co-kr" {
  provider = aws.us-east-1
  domain   = "*.lynlab.co.kr"
  types    = ["AMAZON_ISSUED"]
  statuses = ["ISSUED"]
}
