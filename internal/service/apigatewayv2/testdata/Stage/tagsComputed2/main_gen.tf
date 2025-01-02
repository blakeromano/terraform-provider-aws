# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "null" {}

resource "aws_apigatewayv2_stage" "test" {
  name   = var.rName
  api_id = aws_apigatewayv2_api.test.id

  tags = {
    (var.unknownTagKey) = null_resource.test.id
    (var.knownTagKey)   = var.knownTagValue
  }
}

resource "aws_apigatewayv2_api" "test" {
  name                       = var.rName
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "null_resource" "test" {}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

variable "unknownTagKey" {
  type     = string
  nullable = false
}

variable "knownTagKey" {
  type     = string
  nullable = false
}

variable "knownTagValue" {
  type     = string
  nullable = false
}
