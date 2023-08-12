resource "aws_cloudwatch_log_group" "buslive-api" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.buslive-api.id}/${var.stage}"
  retention_in_days = 14
}
