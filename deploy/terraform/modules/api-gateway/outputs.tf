output "id" {
  value = aws_api_gateway_rest_api.buslive-api.id
}

output "invoke_url" {
  value = aws_api_gateway_deployment.buslive-api.invoke_url
}
