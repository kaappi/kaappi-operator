// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.Kaappi":               schema_pkg_apis_kaappi_v1alpha1_Kaappi(ref),
		"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResource":       schema_pkg_apis_kaappi_v1alpha1_KaappiResource(ref),
		"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResourceSpec":   schema_pkg_apis_kaappi_v1alpha1_KaappiResourceSpec(ref),
		"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResourceStatus": schema_pkg_apis_kaappi_v1alpha1_KaappiResourceStatus(ref),
		"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiSpec":           schema_pkg_apis_kaappi_v1alpha1_KaappiSpec(ref),
		"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiStatus":         schema_pkg_apis_kaappi_v1alpha1_KaappiStatus(ref),
	}
}

func schema_pkg_apis_kaappi_v1alpha1_Kaappi(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kaappi is the Schema for the kaappis API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiSpec", "github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kaappi_v1alpha1_KaappiResource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KaappiResource is the Schema for the kaappiresources API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResourceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResourceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResourceSpec", "github.com/kaappi/kaappi-operator/pkg/apis/kaappi/v1alpha1.KaappiResourceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kaappi_v1alpha1_KaappiResourceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KaappiResourceSpec defines the desired state of KaappiResource",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_kaappi_v1alpha1_KaappiResourceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KaappiResourceStatus defines the observed state of KaappiResource",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_kaappi_v1alpha1_KaappiSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KaappiSpec defines the desired state of Kaappi",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_kaappi_v1alpha1_KaappiStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KaappiStatus defines the observed state of Kaappi",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
