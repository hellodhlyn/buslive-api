data "archive_file" "lambda" {
  type        = "zip"
  source_file = "./../../build/lambda"
  output_path = "./../../build/lambda.zip"
}

resource "aws_lambda_function" "buslive-api" {
  filename         = "./../../build/lambda.zip"
  function_name    = "buslive-api-${var.stage}"
  role             = aws_iam_role.iam_for_lambda.arn
  handler          = "lambda"
  source_code_hash = data.archive_file.lambda.output_base64sha256

  runtime          = "go1.x"
  architectures    = [ "x86_64" ]
  timeout          = 10
  memory_size      = 128

  environment {
    variables = {
      SEOUL_BUS_API_KEY = sensitive(var.seoul_bus_api_key)
    }
  }
}
