// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package webhook

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/gateways/core/webhook.webhook": schema_argo_events_gateways_core_webhook_webhook(ref),
	}
}

func schema_argo_events_gateways_core_webhook_webhook(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "webhook is a general purpose REST API",
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "REST API endpoint",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"method": {
						SchemaProps: spec.SchemaProps{
							Description: "Method is HTTP request method that indicates the desired action to be performed for a given resource. See RFC7231 Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Description: "Port on which HTTP server is listening for incoming events.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"endpoint", "method", "port"},
			},
		},
		Dependencies: []string{},
	}
}