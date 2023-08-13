terraform {
  required_version = "~> 1.5"
}

provider "aws" {
  region = "ap-northeast-2"
}

module "lambda" {
  source = "./modules/lambda"
  stage  = terraform.workspace == "default" ? "prod" : terraform.workspace

  seoul_bus_api_key = sensitive(var.seoul_bus_api_key)
}

module "api-gateway" {
  source = "./modules/api-gateway"
  stage  = terraform.workspace == "default" ? "prod" : terraform.workspace

  lambda_function_name = module.lambda.function_name
  lambda_invoke_arn    = module.lambda.invoke_arn
}
