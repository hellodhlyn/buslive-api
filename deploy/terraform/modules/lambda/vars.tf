variable "stage" {
  type     = string
  nullable = false
}

variable "seoul_bus_api_key" {
  type      = string
  nullable  = false
  sensitive = true
}
