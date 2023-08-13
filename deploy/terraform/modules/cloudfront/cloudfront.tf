data "aws_region" "current" {}

locals {
  origin_id = "api-gateway-${var.api_gateway_id}"
}

resource "aws_cloudfront_cache_policy" "buslive-api" {
  name        = "BusLiveAPIPolicy"
  default_ttl = 3600
  min_ttl     = 600
  max_ttl     = 86400

  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config {
      cookie_behavior = "all"
    }

    headers_config {
      header_behavior = "whitelist"
      headers {
        items = ["Authorization"]
      }
    }

    query_strings_config {
      query_string_behavior = "all"
    }
  }
}

resource "aws_cloudfront_distribution" "buslive-api" {
  depends_on = [
    data.aws_acm_certificate.lynlab-co-kr
  ]

  enabled         = true
  comment         = "BusLive API for ${var.stage} stage"
  is_ipv6_enabled = true
  aliases         = [ "buslive-${var.stage}.lynlab.co.kr" ]

  origin {
    origin_id   = local.origin_id
    domain_name = "${var.api_gateway_id}.execute-api.${data.aws_region.current.name}.amazonaws.com"
    origin_path = "/${var.stage}"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  default_cache_behavior {
    cache_policy_id  = aws_cloudfront_cache_policy.buslive-api.id
    target_origin_id = local.origin_id
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]

    compress               = true
    viewer_protocol_policy = "https-only"
  }

  viewer_certificate {
    acm_certificate_arn      = data.aws_acm_certificate.lynlab-co-kr.arn
    minimum_protocol_version = "TLSv1.2_2021"
    ssl_support_method       = "sni-only"
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["KR", "JP"]
    }
  }
}
