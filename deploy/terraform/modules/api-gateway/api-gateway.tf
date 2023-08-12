resource "aws_api_gateway_rest_api" "buslive-api" {
  name        = "BusLiveAPI"
  description = "BusLive API"
}

resource "aws_api_gateway_resource" "buslive-api" {
  rest_api_id = aws_api_gateway_rest_api.buslive-api.id
  parent_id   = aws_api_gateway_rest_api.buslive-api.root_resource_id
  path_part   = "api"
}

resource "aws_api_gateway_resource" "buslive-api-proxy" {
  rest_api_id = aws_api_gateway_rest_api.buslive-api.id
  parent_id   = aws_api_gateway_resource.buslive-api.id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "buslive-api-proxy" {
  rest_api_id   = aws_api_gateway_rest_api.buslive-api.id
  resource_id   = aws_api_gateway_resource.buslive-api-proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "buslive-lambda" {
  rest_api_id             = aws_api_gateway_rest_api.buslive-api.id
  resource_id             = aws_api_gateway_resource.buslive-api-proxy.id
  http_method             = aws_api_gateway_method.buslive-api-proxy.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${var.lambda_invoke_arn}"
}

resource "aws_api_gateway_deployment" "buslive-api" {
  rest_api_id = aws_api_gateway_rest_api.buslive-api.id
}

resource "aws_api_gateway_stage" "buslive-api" {
  stage_name    = "${var.stage}"
  rest_api_id   = aws_api_gateway_rest_api.buslive-api.id
  deployment_id = aws_api_gateway_deployment.buslive-api.id

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.buslive-api.arn

    # https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html
    format = "$context.identity.sourceIp $context.identity.caller $context.identity.user [$context.requestTime] \"$context.httpMethod $context.resourcePath $context.protocol\" $context.status $context.responseLength $context.requestId"
  }
}
