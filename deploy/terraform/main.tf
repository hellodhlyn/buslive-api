terraform {
  required_version = "~> 1.5"
}

provider "aws" {
  region = "ap-northeast-2"
}

locals {
  stage = terraform.workspace == "default" ? "prod" : terraform.workspace
}

module "lambda" {
  source = "./modules/lambda"
  stage  = local.stage

  seoul_bus_api_key = sensitive(var.seoul_bus_api_key)
}

module "api-gateway" {
  source = "./modules/api-gateway"
  stage  = local.stage

  lambda_function_name = module.lambda.function_name
  lambda_invoke_arn    = module.lambda.invoke_arn
}

module "cloudfront" {
  source = "./modules/cloudfront"
  stage  = local.stage

  api_gateway_id         = module.api-gateway.id
  api_gateway_invoke_url = module.api-gateway.invoke_url
}
