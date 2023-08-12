data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

variable "stage" {
  type     = string
  nullable = false
}

variable "lambda_function_name" {
  type     = string
  nullable = false
}

variable "lambda_invoke_arn" {
  type     = string
  nullable = false
}
